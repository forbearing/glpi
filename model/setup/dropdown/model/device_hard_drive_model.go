package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*DeviceHardDriveModel]()
}

// DeviceHardDriveModel
type DeviceHardDriveModel struct{ model.NameProductNumber }

func (*DeviceHardDriveModel) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*DeviceHardDriveModel]())
}
