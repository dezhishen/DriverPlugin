package exporter

import (
	"context"
	"encoding/json"

	"github.com/OpenListTeam/OpenList/v4/pkg/model"
	"github.com/dezhishen/DriverPlugin/wrap"
)

//go:wasmexport config
func Config(ptr uint32, size uint32) (outPtr uint32, outLen uint32) {
	cfg := wrap.Config()
	data, _ := json.Marshal(cfg)
	return stringToPtr(string(data))
}

//go:wasmexport getStorage
func GetStorage(ptr uint32, size uint32) (outPtr uint32, outLen uint32) {
	storage := wrap.GetStorage()
	data, _ := json.Marshal(storage)
	return stringToPtr(string(data))
}

//go:wasmexport setStorage
func SetStorage(ptr uint32, size uint32) (outPtr uint32, outLen uint32) {
	s := ptrToString(ptr, size)
	var m model.Storage
	_ = json.Unmarshal([]byte(s), &m)
	wrap.SetStorage(m)
	// 返回空字符串
	return stringToPtr("")
}

//go:wasmexport getAddition
func GetAddition(ptr uint32, size uint32) (outPtr uint32, outLen uint32) {
	add := wrap.GetAddition()
	data, _ := json.Marshal(add)
	return stringToPtr(string(data))
}

//go:wasmexport init
func Init(ptr uint32, size uint32) (outPtr uint32, outLen uint32) {
	ctx := context.Background()
	err := wrap.Init(ctx)
	res := struct{ Error string }{Error: ""}
	if err != nil {
		res.Error = err.Error()
	}
	data, _ := json.Marshal(res)
	return stringToPtr(string(data))
}

//go:wasmexport drop
func Drop(ptr uint32, size uint32) (outPtr uint32, outLen uint32) {
	ctx := context.Background()
	err := wrap.Drop(ctx)
	res := struct{ Error string }{Error: ""}
	if err != nil {
		res.Error = err.Error()
	}
	data, _ := json.Marshal(res)
	return stringToPtr(string(data))
}
