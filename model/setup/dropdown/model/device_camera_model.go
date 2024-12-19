package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*DeviceCameraModel]()
}

// DeviceCameraModel
type DeviceCameraModel struct{ model.NameProductNumber }

func (*DeviceCameraModel) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*DeviceCameraModel]())
}
