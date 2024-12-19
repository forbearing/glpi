package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*AutoUpdateSystem]()
}

type AutoUpdateSystem struct{ model.NameEntity }

func (*AutoUpdateSystem) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*AutoUpdateSystem]())
}
