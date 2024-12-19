package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*OperatingSystemKernel]()
}

type OperatingSystemKernel struct{ model.NameEntity }

func (*OperatingSystemKernel) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*OperatingSystemKernel]())
}
