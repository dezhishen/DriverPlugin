package model

var _impl Driver

func Registry(d Driver) {
	_impl = d
}

func Name() string {
	if _impl == nil {
		return "OpenList"
	}
	return _impl.Name()
}
