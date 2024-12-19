package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*DeviceNetworkCardModel]()
}

// DeviceNetworkCardModel
type DeviceNetworkCardModel struct{ model.NameProductNumber }

func (*DeviceNetworkCardModel) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*DeviceNetworkCardModel]())
}
