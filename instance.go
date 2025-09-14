package driverplugin

import (
	"context"
	"fmt"
	"log"

	"github.com/OpenListTeam/OpenList/v4/pkg/driver"
	"github.com/OpenListTeam/OpenList/v4/pkg/model"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

type WasmPluginDriver interface {
	Name(ctx context.Context) string
	driver.Driver
}

type wasmPluginDriver struct {
	ctx       context.Context
	wasmBytes []byte
	runtime   wazero.Runtime
	module    api.Module
	initErr   error
}

func (d *wasmPluginDriver) Init(ctx context.Context) error {
	if d.runtime != nil && d.module != nil {
		return nil // 已初始化
	}
	if ctx == nil {
		ctx = context.Background()
	}
	d.ctx = ctx
	d.runtime = wazero.NewRuntimeWithConfig(d.ctx, wazero.NewRuntimeConfigInterpreter())
	// 注册host函数
	_, err := d.runtime.NewHostModuleBuilder("env").
		NewFunctionBuilder().WithFunc(d.logString).Export("log").
		Instantiate(d.ctx)
	if err != nil {
		d.initErr = err
		return err
	}
	wasi_snapshot_preview1.MustInstantiate(d.ctx, d.runtime)
	mod, err := d.runtime.InstantiateWithConfig(d.ctx, d.wasmBytes, wazero.NewModuleConfig().WithStartFunctions("_initialize"))
	if err != nil {
		d.initErr = err
		return err
	}
	d.module = mod
	//todo need call wasm init function
	return nil
}

func NewWasmPluginDriver(wasmBytes []byte) *wasmPluginDriver {
	return &wasmPluginDriver{
		wasmBytes: wasmBytes,
	}
}

func (d *wasmPluginDriver) Name(ctx context.Context) string {
	d.Init(ctx)
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

func (d *wasmPluginDriver) logString(_ context.Context, m api.Module, offset, byteCount uint32) {
	buf, ok := m.Memory().Read(offset, byteCount)
	if !ok {
		log.Panicf("Memory.Read(%d, %d) out of range", offset, byteCount)
	}
	fmt.Println(string(buf))
}

func (d *wasmPluginDriver) Close() {
	if d.runtime != nil {
		if err := d.runtime.Close(d.ctx); err != nil {
			log.Println("Failed to close runtime:", err)
		}
		d.runtime = nil
		d.module = nil
	}
}

// 实现 driver.Driver 接口的其他方法
var _ WasmPluginDriver = (*wasmPluginDriver)(nil)

// Meta implements the Driver interface.

// Config() Config
// 	// GetStorage just get raw storage, no need to implement, because model.Storage have implemented
// GetStorage() *model.Storage
// SetStorage(model.Storage)
// // GetAddition Additional is used for unmarshal of JSON, so need return pointer
// GetAddition() Additional
// // Init If already initialized, drop first
// Init(ctx context.Context) error
// Drop(ctx context.Context) error

func (d *wasmPluginDriver) Config() driver.Config {
	panic("not implemented")
}

func (d *wasmPluginDriver) GetStorage() *model.Storage {
	return nil
}

func (d *wasmPluginDriver) SetStorage(storage model.Storage) {

}

func (d *wasmPluginDriver) GetAddition() driver.Additional {
	return nil
}

func (d *wasmPluginDriver) Drop(ctx context.Context) error {
	return nil
}

// // List files in the path
// // if identify files by path, need to set ID with path,like path.Join(dir.GetID(), obj.GetName())
// // if identify files by id, need to set ID with corresponding id
// List(ctx context.Context, dir model.Obj, args model.ListArgs) ([]model.Obj, error)
// // Link get url/filepath/reader of file
// Link(ctx context.Context, file model.Obj, args model.LinkArgs) (*model.Link, error)

func (d *wasmPluginDriver) List(ctx context.Context, dir model.Obj, args model.ListArgs) ([]model.Obj, error) {
	panic("not implemented")
}

func (d *wasmPluginDriver) Link(ctx context.Context, file model.Obj, args model.LinkArgs) (*model.Link, error) {
	panic("not implemented")
}
