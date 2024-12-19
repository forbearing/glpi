package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*DeviceDriveModel]()
}

// DeviceDriveModel
type DeviceDriveModel struct{ model.NameProductNumber }

func (*DeviceDriveModel) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*DeviceDriveModel]())
}
