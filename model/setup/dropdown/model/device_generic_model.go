package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*DeviceGenericModel]()
}

// DeviceGenericModel
type DeviceGenericModel struct{ model.NameProductNumber }

func (*DeviceGenericModel) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*DeviceGenericModel]())
}
