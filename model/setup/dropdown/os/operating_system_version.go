package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*OperatingSystemVersion]()
}

type OperatingSystemVersion struct{ model.NameEntity }

func (*OperatingSystemVersion) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*OperatingSystemVersion]())
}
