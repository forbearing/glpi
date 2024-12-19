package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*OperatingSystemArchitecture]()
}

type OperatingSystemArchitecture struct{ model.NameEntity }

func (*OperatingSystemArchitecture) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*OperatingSystemArchitecture]())
}
