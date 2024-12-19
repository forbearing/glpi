package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*DevicePowerSupplyModel]()
}

// DevicePowerSupplyModel
type DevicePowerSupplyModel struct{ model.NameProductNumber }

func (*DevicePowerSupplyModel) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*DevicePowerSupplyModel]())
}
