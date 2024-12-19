package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*OperatingSystemKernelVersion]()
}

type OperatingSystemKernelVersion struct {
	pkgmodel.Base

	Name                    string  `json:"name"`
	OperatingSystemKernelID *string `json:"operating_system_kernel_id"`
}

func (*OperatingSystemKernelVersion) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*OperatingSystemKernelVersion]())
}
