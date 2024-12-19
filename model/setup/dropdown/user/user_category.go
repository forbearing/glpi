package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*UserCategory]()
}

// UserCategory
type UserCategory struct{ model.NameEntity }

func (*UserCategory) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*UserCategory]())
}
