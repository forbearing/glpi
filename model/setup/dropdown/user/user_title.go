package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*UserTitle]()
}

// UserTitle
type UserTitle struct{ model.NameEntity }

func (*UserTitle) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*UserTitle]())
}
