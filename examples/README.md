# unimplementedWasmPluginDriver
## 编译
```bash
tinygo build -buildmode=c-shared -target=wasip1 -o wasm/unimplemented.wasm wasm/unimplemented.go
```
## 使用
```bash
go run main.go
```