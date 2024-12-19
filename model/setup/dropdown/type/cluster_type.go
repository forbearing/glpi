package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*ClusterType]()
}

type ClusterType struct{ model.NameEntity }

func (*ClusterType) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*ClusterType]())
}
