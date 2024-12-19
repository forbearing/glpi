package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*DeviceControlModel]()
}

// DeviceControlModel
type DeviceControlModel struct{ model.NameProductNumber }

func (*DeviceControlModel) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*DeviceControlModel]())
}
