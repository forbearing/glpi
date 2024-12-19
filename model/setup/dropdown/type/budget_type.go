package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*BudgetType]()
}

type BudgetType struct{ model.NameEntity }

func (*BudgetType) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*BudgetType]())
}
