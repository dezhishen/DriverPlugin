package model

import (
	"context"
	"log"
)

var _impl Driver

func Registry(d Driver) {
	_impl = d
}

func Name(ctx context.Context) string {
	if _impl == nil {
		log.Panicln("Driver plugin not registered")
	}
	return _impl.Name(ctx)
}
