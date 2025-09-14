package wasm

import (
	"github.com/OpenListTeam/OpenList/v4/driverplugin/local"
	"github.com/dezhishen/DriverPlugin/wrap"
)

func init() {
	wrap.Registry(&local.Local{})
}
