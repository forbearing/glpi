package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*OperatingSystemEdition]()
}

type OperatingSystemEdition struct{ model.NameEntity }

func (*OperatingSystemEdition) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*OperatingSystemEdition]())
}
