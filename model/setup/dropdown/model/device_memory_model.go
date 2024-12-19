package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*DeviceMemoryModel]()
}

// DeviceMemoryModel
type DeviceMemoryModel struct{ model.NameProductNumber }

func (*DeviceMemoryModel) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*DeviceMemoryModel]())
}
