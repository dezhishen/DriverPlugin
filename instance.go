package driverplugin

import (
	"context"
	"fmt"
	"log"

	"github.com/OpenListTeam/OpenList/v5/layers/file"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

type WasmPluginDriver interface {
	file.HostFileServer
	Name(ctx context.Context) string
}

type wasmPluginDriver struct {
	ctx       context.Context
	wasmBytes []byte
	runtime   wazero.Runtime
	module    api.Module
	initErr   error
}

var _ WasmPluginDriver = (*wasmPluginDriver)(nil)

func (d *wasmPluginDriver) Close() {
	if d.runtime != nil {
		if err := d.runtime.Close(d.ctx); err != nil {
			log.Println("Failed to close runtime:", err)
		}
		d.runtime = nil
		d.module = nil
	}
}
func (d *wasmPluginDriver) Init() {
	if d.runtime != nil && d.module != nil {
		return // 已初始化
	}
	d.ctx = context.Background()
	d.runtime = wazero.NewRuntimeWithConfig(d.ctx, wazero.NewRuntimeConfigInterpreter())
	// 注册host函数
	_, err := d.runtime.NewHostModuleBuilder("env").
		NewFunctionBuilder().WithFunc(d.logString).Export("log").
		Instantiate(d.ctx)
	if err != nil {
		d.initErr = err
		return
	}
	wasi_snapshot_preview1.MustInstantiate(d.ctx, d.runtime)
	mod, err := d.runtime.InstantiateWithConfig(d.ctx, d.wasmBytes, wazero.NewModuleConfig().WithStartFunctions("_initialize"))
	if err != nil {
		d.initErr = err
		return
	}
	d.module = mod
}

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

func (d *wasmPluginDriver) Name(ctx context.Context) string {
	d.Init()
	if d.initErr != nil {
		log.Panicln(d.initErr)
	}
	mod := d.module
	name := mod.ExportedFunction("name")
	malloc := mod.ExportedFunction("malloc")
	free := mod.ExportedFunction("free")
	results, err := malloc.Call(ctx, 128)
	if err != nil {
		log.Panicln(err)
	}
	namePtr := results[0]
	nameLen := uint64(128)
	defer free.Call(ctx, namePtr)
	results, err = name.Call(ctx, namePtr, nameLen)
	if err != nil {
		log.Panicln(err)
	}
	nameLen = results[0]
	nameBytes, ok := mod.Memory().Read(uint32(namePtr), uint32(nameLen))
	if !ok {
		log.Panicf("Memory.Read(%d, %d) out of range", namePtr, nameLen)
	}
	return string(nameBytes)
}

func NewWasmPluginDriver(wasmBytes []byte) *wasmPluginDriver {
	return &wasmPluginDriver{
		wasmBytes: wasmBytes,
	}
}

func (d *wasmPluginDriver) logString(_ context.Context, m api.Module, offset, byteCount uint32) {
	buf, ok := m.Memory().Read(offset, byteCount)
	if !ok {
		log.Panicf("Memory.Read(%d, %d) out of range", offset, byteCount)
	}
	fmt.Println(string(buf))
}
