package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*DevicePciModel]()
}

// DevicePciModel
type DevicePciModel struct{ model.NameProductNumber }

func (*DevicePciModel) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*DevicePciModel]())
}
