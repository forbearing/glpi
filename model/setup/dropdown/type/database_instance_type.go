package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*DatabaseInstanceType]()
}

type DatabaseInstanceType struct{ model.NameEntity }

func (*DatabaseInstanceType) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*DatabaseInstanceType]())
}
