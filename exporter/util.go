package exporter

// #include <stdlib.h>
import "C"
import (
	"encoding/json"
	"unsafe"
)

// structToPtr 将任意结构体编码为 JSON，并返回可用于 WebAssembly 的指针和长度。
// 注意：返回的指针直接指向 Go 堆上的 slice，调用方需保证生命周期（建议仅用于 wasm host/guest互操作）。
func structToPtr(v interface{}) (ptr uint32, size uint32) {
	data, _ := json.Marshal(v)
	if len(data) == 0 {
		return 0, 0
	}
	// 返回指针和长度
	return uint32(uintptr(unsafe.Pointer(&data[0]))), uint32(len(data))
}

// ptrToStruct 从 WebAssembly 内存中的指针和长度读取数据，并解码为结构体。
// v 必须为指向目标结构体的指针。
func ptrToStruct(ptr uint32, size uint32, v interface{}) error {
	if size == 0 {
		return nil
	}
	data := unsafe.Slice((*byte)(unsafe.Pointer(uintptr(ptr))), size)
	return json.Unmarshal(data, v)
}

// stringToPtr 将字符串转换为 WebAssembly 可用的指针和长度。
// 注意：返回的指针直接指向 Go 堆上的 slice，调用方需保证生命周期（建议仅用于 wasm host/guest互操作）。
func stringToPtr(s string) (ptr uint32, size uint32) {
	if len(s) == 0 {
		return 0, 0
	}
	// 返回指针和长度
	return uint32(uintptr(unsafe.Pointer(&[]byte(s)[0]))), uint32(len(s))
}
