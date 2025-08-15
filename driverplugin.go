package driverplugin

import (
	"os"
	"path/filepath"

	"github.com/dezhishen/DriverPlugin/model"
)

// LoadPlugins loads all driver plugins from the specified path.
func LoadPlugins(path string) ([]model.Driver, error) {
	// read all *.wasm files in the directory
	// scan the directory for .wasm files
	matches, err := filepath.Glob(filepath.Join(path, "*.wasm"))
	if err != nil {
		return nil, err
	}
	drivers := make([]model.Driver, 0, len(matches))
	for _, match := range matches {
		driver, err := newDriverPlugin(match)
		if err != nil {
			return nil, err
		}
		drivers = append(drivers, driver)
	}
	return drivers, nil
}

func newDriverPlugin(path string) (model.Driver, error) {
	// create a new driver plugin
	wasmBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return NewWasmPluginDriver(wasmBytes), err
}
