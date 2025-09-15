package exporter

import (
	"runtime"

	"github.com/dezhishen/DriverPlugin/wrap"
)

//go:wasmexport name
func Name() (uint32, uint32) {
	driverName := wrap.Name()
	return stringToPtr(driverName)
}

// Log a message to the console using _log.
func Log(message string) {
	ptr, size := stringToPtr(message)
	_log(ptr, size)
	runtime.KeepAlive(message) // keep message alive until ptr is no longer needed.
}

// _log is a WebAssembly import which prints a string (linear memory offset,
// byteCount) to the console.
//
//go:wasmimport env log
func _log(ptr, size uint32)
