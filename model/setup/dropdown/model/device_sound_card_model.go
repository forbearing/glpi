package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*DeviceSoundCardModel]()
}

// DeviceSoundCardModel
type DeviceSoundCardModel struct{ model.NameProductNumber }

func (*DeviceSoundCardModel) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*DeviceSoundCardModel]())
}
