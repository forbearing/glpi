package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*Manufacturer]()
}

// Manufacturer 制造商信息
type Manufacturer struct {
	pkgmodel.Base

	// 基础信息
	Name *string `json:"name"` // 制造商名称
}

func (*Manufacturer) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*Manufacturer]())
}
