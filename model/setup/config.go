package model

import (
	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*Config]()
}

type Config struct {
	model.Name
}
