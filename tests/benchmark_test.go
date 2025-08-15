package tests

import (
	"context"
	"testing"

	driverplugin "github.com/dezhishen/DriverPlugin"
)

func BenchmarkCallPlugins(b *testing.B) {
	// compile the wasm files first if needed
	drivers, err := driverplugin.LoadPlugins("../examples/wasm")
	if err != nil {
		b.Fatal(err)
	}
	// use the first driver for benchmarking
	if len(drivers) == 0 {
		b.Fatal("no drivers loaded")
	}
	driver := drivers[0]
	ctx := context.Background()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		name := driver.Name(ctx)
		if name != "unimplemented" {
			b.Fatal("driver name is empty")
		}
		ctx.Done()
	}
}
