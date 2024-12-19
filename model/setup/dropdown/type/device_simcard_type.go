package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*DeviceSimcardType]()
}

type DeviceSimcardType struct{ model.NameEntity }

func (*DeviceSimcardType) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*DeviceSimcardType]())
}
