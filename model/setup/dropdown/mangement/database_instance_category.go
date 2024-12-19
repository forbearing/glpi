package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*DatabaseInstanceCategory]()
}

type DatabaseInstanceCategory struct{ model.NameEntity }

func (*DatabaseInstanceCategory) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*DatabaseInstanceCategory]())
}
