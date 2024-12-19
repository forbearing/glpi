package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*DeviceProcessorModel]()
}

// DeviceProcessorModel
type DeviceProcessorModel struct{ model.NameProductNumber }

func (*DeviceProcessorModel) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*DeviceProcessorModel]())
}
