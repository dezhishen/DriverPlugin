package main

import (
	"github.com/dezhishen/DriverPlugin/examples/local"
	"github.com/dezhishen/DriverPlugin/wrap"
)

func init() {
	wrap.Registry(&local.Local{})
}
