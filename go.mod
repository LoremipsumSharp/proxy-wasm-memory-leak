module github.com/LoremipsumSharp/proxy-wasm-memory-leak

go 1.19

require (
	github.com/VictoriaMetrics/fastcache v0.0.0-00010101000000-000000000000
	github.com/google/uuid v1.3.0
	github.com/tetratelabs/proxy-wasm-go-sdk v0.22.0
	github.com/wasilibs/nottinygc v0.4.0
)

require (
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/magefile/mage v1.14.0 // indirect
)

replace github.com/VictoriaMetrics/fastcache v0.0.0-00010101000000-000000000000 => github.com/LoremipsumSharp/fastcache v0.0.0-20230714190828-f49a825cb86c
