package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*DeviceGraphicCardModel]()
}

// DeviceGraphicCardModel
type DeviceGraphicCardModel struct{ model.NameProductNumber }

func (*DeviceGraphicCardModel) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*DeviceGraphicCardModel]())
}
