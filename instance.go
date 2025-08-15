package driverplugin

import (
	"context"

	"github.com/OpenListTeam/OpenList/v5/layers/file"
	"github.com/dezhishen/DriverPlugin/model"
)

type wasmPluginDriver struct {
	wasmBytes []byte
}

var _ model.Driver = (*wasmPluginDriver)(nil)

// CopyFile implements model.Driver.
func (d *wasmPluginDriver) CopyFile(ctx context.Context, sources []string, targets []string) ([]*file.BackFileAction, error) {
	panic("unimplemented")
}

// Download implements model.Driver.
func (d *wasmPluginDriver) Download(ctx context.Context, path []string, opt *file.DownloadOption) ([]*file.LinkFileObject, error) {
	panic("unimplemented")
}

// FindFile implements model.Driver.
func (d *wasmPluginDriver) FindFile(ctx context.Context, path []string, opt *file.FindFileOption) ([]*file.HostFileObject, error) {
	panic("unimplemented")
}

// KillFile implements model.Driver.
func (d *wasmPluginDriver) KillFile(ctx context.Context, path []string, opt *file.KillFileOption) ([]*file.BackFileAction, error) {
	panic("unimplemented")
}

// ListFile implements model.Driver.
func (d *wasmPluginDriver) ListFile(ctx context.Context, path []string, opt *file.ListFileOption) ([]*file.HostFileObject, error) {
	panic("unimplemented")
}

// MakeFile implements model.Driver.
func (d *wasmPluginDriver) MakeFile(ctx context.Context, path []string, opt *file.MakeFileOption) ([]*file.BackFileAction, error) {
	panic("unimplemented")
}

// MakePath implements model.Driver.
func (d *wasmPluginDriver) MakePath(ctx context.Context, path []string, opt *file.MakeFileOption) ([]*file.BackFileAction, error) {
	panic("unimplemented")
}

// MoveFile implements model.Driver.
func (d *wasmPluginDriver) MoveFile(ctx context.Context, sources []string, targets []string) ([]*file.BackFileAction, error) {
	panic("unimplemented")
}

// NameFile implements model.Driver.
func (d *wasmPluginDriver) NameFile(ctx context.Context, sources []string, targets []string) ([]*file.BackFileAction, error) {
	panic("unimplemented")
}

// Uploader implements model.Driver.
func (d *wasmPluginDriver) Uploader(ctx context.Context, path []string, opt *file.UploaderOption) ([]*file.BackFileAction, error) {
	panic("unimplemented")
}

func (d *wasmPluginDriver) Name() string {
	return ""
}

func NewWasmPluginDriver(wasmBytes []byte) *wasmPluginDriver {
	return &wasmPluginDriver{
		wasmBytes: wasmBytes,
	}
}
