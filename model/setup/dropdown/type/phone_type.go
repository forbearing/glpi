package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*PhoneType]()
}

type PhoneType struct{ model.NameEntity }

func (*PhoneType) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*PhoneType]())
}
