package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*PassiveDeviceEquipmentType]()
}

type PassiveDeviceEquipmentType struct{ model.NameEntity }

func (*PassiveDeviceEquipmentType) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*PassiveDeviceEquipmentType]())
}
