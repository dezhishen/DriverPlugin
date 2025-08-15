package model

import (
	"github.com/OpenListTeam/OpenList/v5/layers/file"
)

type Driver interface {
	file.HostFileServer
	Name() string
}
