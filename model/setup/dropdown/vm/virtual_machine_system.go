package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*VirtualMachineSystem]()
}

// VirtualMachineSystem
type VirtualMachineSystem struct{ model.NameEntity }

func (*VirtualMachineSystem) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*VirtualMachineSystem]())
}
