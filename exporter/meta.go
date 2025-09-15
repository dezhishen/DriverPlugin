package exporter

import (
	"context"

	"github.com/OpenListTeam/OpenList/v4/pkg/model"
	"github.com/dezhishen/DriverPlugin/wrap"
)

//go:wasmexport config
func Config(ptr uint32, size uint32) (outPtr uint32, outLen uint32) {
	return structToPtr(wrap.Config())
}

//go:wasmexport getStorage
func GetStorage(ptr uint32, size uint32) (outPtr uint32, outLen uint32) {
	return structToPtr(wrap.GetStorage())
}

//go:wasmexport setStorage
func SetStorage(ptr uint32, size uint32) (outPtr uint32, outLen uint32) {
	var m model.Storage
	err := ptrToStruct(ptr, size, &m)
	if err != nil {
		// 返回错误信息
		res := struct{ Error string }{Error: err.Error()}
		return structToPtr(res)
	}
	wrap.SetStorage(m)
	// 返回空字符串
	return structToPtr(struct{ Error string }{Error: ""})
}

//go:wasmexport getAddition
func GetAddition(ptr uint32, size uint32) (outPtr uint32, outLen uint32) {
	return structToPtr(wrap.GetAddition())
}

//go:wasmexport init
func Init(ptr uint32, size uint32) (outPtr uint32, outLen uint32) {
	ctx := context.Background()
	err := wrap.Init(ctx)
	res := struct{ Error string }{Error: ""}
	if err != nil {
		res.Error = err.Error()
	}
	return structToPtr(res)
}

//go:wasmexport drop
func Drop(ptr uint32, size uint32) (outPtr uint32, outLen uint32) {
	ctx := context.Background()
	err := wrap.Drop(ctx)
	res := struct{ Error string }{Error: ""}
	if err != nil {
		res.Error = err.Error()
	}
	return structToPtr(res)
}
