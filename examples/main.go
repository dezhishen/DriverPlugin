package main

import (
	"context"

	driverplugin "github.com/dezhishen/DriverPlugin"
)

func main() {
	drives, err := driverplugin.LoadPlugins("./wasm")
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	for _, driver := range drives {
		name := driver.Name(ctx)
		if name == "" {
			panic("driver name is empty")
		}
		println("Loaded driver:", name)
	}
	// You can now use the loaded drivers as needed
	// For example, you can call methods on the driver instances
	// driver.CopyFile(...), driver.Download(...), etc.
	// This is just a placeholder main function to demonstrate loading plugins
	// and you can expand it based on your application's requirements.
	// Note: Ensure that the WASM files in the ./wasm directory implement the required
	// methods defined in the model.DriverPlugin interface.
}
