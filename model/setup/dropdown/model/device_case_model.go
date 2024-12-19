package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*DeviceCaseModel]()
}

// DeviceCaseModel
type DeviceCaseModel struct{ model.NameProductNumber }

func (*DeviceCaseModel) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*DeviceCaseModel]())
}
