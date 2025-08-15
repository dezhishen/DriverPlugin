package model

import (
	"context"

	"github.com/OpenListTeam/OpenList/v5/layers/file"
)

type Driver interface {
	file.HostFileServer
	Name(ctx context.Context) string
}
