package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*DeviceMotherBoardModel]()
}

// DeviceMotherBoardModel
type DeviceMotherBoardModel struct{ model.NameProductNumber }

func (*DeviceMotherBoardModel) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*DeviceMotherBoardModel]())
}
