package main

import (
	"github.com/OpenListTeam/OpenList/v5/layers/file"
	_ "github.com/dezhishen/DriverPlugin/expoter"
	"github.com/dezhishen/DriverPlugin/model"
)

func main() {}

type unimplementedWasmPluginDriver struct{}

func init() {
	model.Registry(&unimplementedWasmPluginDriver{})
}

func (d *unimplementedWasmPluginDriver) Name() string {
	return "unimplemented"
}

func (d *unimplementedWasmPluginDriver) CopyFile(sources []string, targets []string) ([]*file.BackFileAction, error) {
	panic("unimplemented")
}

func (d *unimplementedWasmPluginDriver) MoveFile(sources []string, targets []string) ([]*file.BackFileAction, error) {
	panic("unimplemented")
}

func (d *unimplementedWasmPluginDriver) ListFile(path []string, opt *file.ListFileOption) ([]*file.HostFileObject, error) {
	panic("unimplemented")
}

func (d *unimplementedWasmPluginDriver) FindFile(path []string, opt *file.FindFileOption) ([]*file.HostFileObject, error) {
	panic("unimplemented")
}

func (d *unimplementedWasmPluginDriver) Download(path []string, opt *file.DownloadOption) ([]*file.LinkFileObject, error) {
	panic("unimplemented")
}

func (d *unimplementedWasmPluginDriver) Uploader(path []string, opt *file.UploaderOption) ([]*file.BackFileAction, error) {
	panic("unimplemented")
}

func (d *unimplementedWasmPluginDriver) RemoveFile(path []string, opt *file.KillFileOption) ([]*file.BackFileAction, error) {
	panic("unimplemented")
}

func (d *unimplementedWasmPluginDriver) MakeFile(path []string, opt *file.MakeFileOption) ([]*file.BackFileAction, error) {
	panic("unimplemented")
}

func (d *unimplementedWasmPluginDriver) MakePath(path []string, opt *file.MakeFileOption) ([]*file.BackFileAction, error) {
	panic("unimplemented")
}
