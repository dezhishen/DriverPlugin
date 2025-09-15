package exporter

import (
	"context"

	"github.com/OpenListTeam/OpenList/v4/pkg/model"
	"github.com/dezhishen/DriverPlugin/wrap"
)

//go:wasmexport list
func List(ptr uint32, size uint32) (outPtr uint32, outLen uint32) {
	var args struct {
		Dir  model.Obj      `json:"dir"`
		Args model.ListArgs `json:"args"`
	}
	err := ptrToStruct(ptr, size, &args)
	res := struct {
		Objs  interface{} `json:"objs"`
		Error string      `json:"error"`
	}{}
	if err != nil {
		res.Error = err.Error()
		return structToPtr(res)
	}
	objs, err := wrap.List(context.Background(), args.Dir, args.Args)
	if err != nil {
		res.Error = err.Error()
		return structToPtr(res)
	}
	res.Objs = objs
	return structToPtr(res)
}

//go:wasmexport link
func Link(ptr uint32, size uint32) (outPtr uint32, outLen uint32) {
	var args struct {
		File     model.Obj      `json:"file"`
		LinkArgs model.LinkArgs `json:"args"`
	}
	err := ptrToStruct(ptr, size, &args)
	if err != nil {
		res := struct {
			Link  *model.Link `json:"link"`
			Error string      `json:"error"`
		}{Link: nil, Error: err.Error()}
		return structToPtr(res)
	}
	link, err := wrap.Link(context.Background(), args.File, args.LinkArgs)
	res := struct {
		Link  *model.Link `json:"link"`
		Error string      `json:"error"`
	}{Link: link, Error: ""}
	if err != nil {
		res.Error = err.Error()
	}
	return structToPtr(res)
}
