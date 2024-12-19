package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*VirtualMachineType]()
}

type VirtualMachineType struct{ model.NameEntity }

func (*VirtualMachineType) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*VirtualMachineType]())
}
