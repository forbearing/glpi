package model

import (
	"fmt"

	pkgmodel "github.com/forbearing/golib/model"

	"glpi/model"
)

func init() {
	pkgmodel.Register[*ApplianceType]()
}

type ApplianceType struct{ model.Name }

func (*ApplianceType) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*ApplianceType]())
}
