package main

import (
	"runtime"

	"github.com/VictoriaMetrics/fastcache"
	"github.com/google/uuid"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
	_ "github.com/wasilibs/nottinygc"
)

var (
	cache *fastcache.Cache = fastcache.New(32 * 1024 * 1024)
)

type (
	vmContext     struct{}
	pluginContext struct {
		types.DefaultPluginContext
	}

	httpContext struct {
		types.DefaultHttpContext
	}

	Foo struct {
		Bar string
	}
)

func (*vmContext) OnVMStart(vmConfigurationSize int) types.OnVMStartStatus {
	proxywasm.LogCritical("successfully started VM.")
	return types.OnVMStartStatusOK
}

func (ctx *pluginContext) NewHttpContext(contextID uint32) types.HttpContext {
	return &httpContext{}
}

func (*vmContext) NewPluginContext(contextID uint32) types.PluginContext {
	return &pluginContext{}
}
func (*httpContext) OnHttpRequestHeaders(numHeaders int, endOfStream bool) types.Action {
	cache.Set([]byte(uuid.Must(uuid.NewRandom()).String()), []byte(uuid.Must(uuid.NewRandom()).String()))
	logMemStats()

	return types.ActionContinue
}

func (*httpContext) OnHttpResponseHeaders(numHeaders int, endOfStream bool) types.Action {
	// proxywasm.AddHttpResponseHeader("memory-usage", fmt.Sprint(memoryUsage()/1024/1024))
	return types.ActionContinue
}

// Override types.DefaultPluginContext.
func (ctx *pluginContext) OnPluginStart(pluginConfigurationSize int) types.OnPluginStartStatus {
	return types.OnPluginStartStatusOK
}

func main() {
	proxywasm.SetVMContext(&vmContext{})
}

//export sched_yield
func sched_yield() int32 {
	return 0
}

func logMemStats() {
	ms := runtime.MemStats{}
	runtime.ReadMemStats(&ms)
	proxywasm.LogWarnf(
		"Sys: %d, HeapSys: %d, HeapIdle: %d, HeapReleased: %d, TotalAlloc: %d",
		ms.Sys,
		ms.HeapSys,
		ms.HeapIdle,
		ms.HeapReleased,
		ms.TotalAlloc)
}
