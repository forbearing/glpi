package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*PhonePowerSupplyType]()
}

type PhonePowerSupplyType struct{ model.NameEntity }

func (*PhonePowerSupplyType) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*PhonePowerSupplyType]())
}
