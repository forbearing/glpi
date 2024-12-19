package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*DeviceSensorModel]()
}

// DeviceSensorModel
type DeviceSensorModel struct{ model.NameProductNumber }

func (*DeviceSensorModel) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*DeviceSensorModel]())
}
