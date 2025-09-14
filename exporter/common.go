package exporter

import (
	"runtime"
	"unsafe"

	"github.com/dezhishen/DriverPlugin/wrap"
)

//go:wasmexport name
func Name(ptr uint32, length uint32) uint32 {
	s := unsafe.Slice((*byte)(unsafe.Pointer(uintptr(ptr))), length)
	driverName := wrap.Name()
	copy(s, driverName)
	return uint32(len(driverName))
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
