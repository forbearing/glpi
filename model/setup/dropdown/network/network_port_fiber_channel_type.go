package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*NetworkPortFiberChannelType]()
}

type NetworkPortFiberChannelType struct {
	model.NameEntity
}

func (*NetworkPortFiberChannelType) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*NetworkPortFiberChannelType]())
}
