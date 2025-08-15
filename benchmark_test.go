package driverplugin

import (
	"context"
	"testing"
)

func BenchmarkCallWithInitPlugin(b *testing.B) {
	// compile the wasm files first if needed
	drivers, err := LoadPlugins("./examples/wasm")
	if err != nil {
		b.Fatal(err)
	}
	// use the first driver for benchmarking
	if len(drivers) == 0 {
		b.Fatal("no drivers loaded")
	}
	driver := drivers[0]
	ctx := context.Background()
	// 强转为 WasmPluginDriver
	wasmDriver, ok := driver.(*wasmPluginDriver)
	if !ok {
		b.Fatal("driver is not a WasmPluginDriver")
	}
	// 初始化插件
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		wasmDriver.runtime = nil // 重置 runtime 以便下次测试
		b.StartTimer()
		wasmDriver.Name(ctx)
	}
}
func BenchmarkCallWithOutInitPlugin(b *testing.B) {
	// compile the wasm files first if needed
	drivers, err := LoadPlugins("./examples/wasm")
	if err != nil {
		b.Fatal(err)
	}
	// use the first driver for benchmarking
	if len(drivers) == 0 {
		b.Fatal("no drivers loaded")
	}
	driver := drivers[0]
	ctx := context.Background()
	wasmDriver, ok := driver.(*wasmPluginDriver)
	if !ok {
		b.Fatal("driver is not a WasmPluginDriver")
	}
	wasmDriver.Init(ctx)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		name := wasmDriver.Name(ctx)
		if name != "unimplemented" {
			b.Fatal("driver name is empty")
		}
		ctx.Done()
	}
}
