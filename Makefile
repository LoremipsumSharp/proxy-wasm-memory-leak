build:
	tinygo build -o main.wasm -gc=custom -tags=custommalloc -target=wasi -scheduler=none main.go

