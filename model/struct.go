package model

import (
	"log"
)

var _impl DriverPlugin

func Registry(d DriverPlugin) {
	_impl = d
}

func Name() string {
	if _impl == nil {
		log.Panicln("Driver plugin not registered")
	}
	return _impl.Name()
}
