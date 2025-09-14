package wrap

import (
	"context"

	"github.com/OpenListTeam/OpenList/v4/pkg/driver"
	"github.com/OpenListTeam/OpenList/v4/pkg/model"
)

var _impl driver.Driver

func Registry(d driver.Driver) {
	_impl = d
}

func Name() string {
	return _impl.Config().Name
}

func Config() driver.Config {
	return _impl.Config()
}

// GetStorage just get raw storage, no need to implement, because model.Storage have implemented
func GetStorage() *model.Storage {
	return _impl.GetStorage()
}
func SetStorage(m model.Storage) {
	_impl.SetStorage(m)
}

// GetAddition Additional is used for unmarshal of JSON, so need return pointer
func GetAddition() driver.Additional {
	return _impl.GetAddition()
}

// Init If already initialized, drop first
func Init(ctx context.Context) error {
	return _impl.Init(ctx)
}

func Drop(ctx context.Context) error {
	return _impl.Drop(ctx)
}
func List(ctx context.Context, dir model.Obj, args model.ListArgs) ([]model.Obj, error) {
	return _impl.List(ctx, dir, args)
}

// Link get url/filepath/reader of file
func Link(ctx context.Context, file model.Obj, args model.LinkArgs) (*model.Link, error) {
	return _impl.Link(ctx, file, args)
}
