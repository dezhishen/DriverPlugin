package local

import (
	"context"
	"errors"

	"github.com/OpenListTeam/OpenList/v4/pkg/model"
	driverplugin "github.com/dezhishen/DriverPlugin"
)

func main() {
	err := load()
	if err != nil {
		panic(err)
	}
	// 获取 local 插件
	localPlugin, exists := allPlugins["local"]
	if !exists {
		panic("local plugin not found")
	}
	// 初始化插件
	err = localPlugin.Init(context.Background())
	if err != nil {
		panic(err)
	}
	// 使用插件的 List 方法
	var dir model.Obj
	// 设置 ListArgs，如果有需要的话
	var args model.ListArgs
	args = model.ListArgs{
		// 设置参数
	}
	// 调用 List 方法
	items, err := localPlugin.List(context.Background(), dir, args)
	if err != nil {
		panic(err)
	}
}

var allPlugins = make(map[string]driverplugin.WasmPluginDriver)

func load() error {
	plugins, err := driverplugin.LoadPlugins("./plugins")
	if err != nil {
		return err
	}
	if len(plugins) == 0 {
		return errors.New("no plugins found")
	}
	for _, p := range plugins {
		if p == nil {
			continue
		}
		name := p.Name(context.Background())
		if name == "" {
			continue
		}
		allPlugins[name] = p
	}
	return nil
}
