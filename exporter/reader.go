package exporter

import (
	"context"
	"encoding/json"

	"github.com/OpenListTeam/OpenList/v4/pkg/model"
	"github.com/dezhishen/DriverPlugin/wrap"
)

//go:wasmexport list
func List(ptr uint32, size uint32) (outPtr uint32, outLen uint32) {
	s := ptrToString(ptr, size)
	var args struct {
		Dir  model.Obj      `json:"dir"`
		Args model.ListArgs `json:"args"`
	}
	_ = json.Unmarshal([]byte(s), &args)
	objs, err := wrap.List(context.Background(), args.Dir, args.Args)
	res := struct {
		Objs  interface{} `json:"objs"`
		Error string      `json:"error"`
	}{Objs: objs, Error: ""}
	if err != nil {
		res.Error = err.Error()
	}
	data, _ := json.Marshal(res)
	return stringToPtr(string(data))
}

//go:wasmexport link
func Link(ptr uint32, size uint32) (outPtr uint32, outLen uint32) {
	s := ptrToString(ptr, size)
	var args struct {
		File     model.Obj      `json:"file"`
		LinkArgs model.LinkArgs `json:"args"`
	}
	_ = json.Unmarshal([]byte(s), &args)
	link, err := wrap.Link(context.Background(), args.File, args.LinkArgs)
	res := struct {
		Link  *model.Link `json:"link"`
		Error string      `json:"error"`
	}{Link: link, Error: ""}
	if err != nil {
		res.Error = err.Error()
	}
	data, _ := json.Marshal(res)
	return stringToPtr(string(data))
}
