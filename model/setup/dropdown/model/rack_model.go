package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*RackModel]()
}

// RackModel 机柜型号
type RackModel struct {
	pkgmodel.Base

	Name          *string `json:"name"`           // 名称
	ProductNumber *string `json:"product_number"` // 产品编号
	Picture       *string `json:"picture"`        // 图片
}

func (*RackModel) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*RackModel]())
}
