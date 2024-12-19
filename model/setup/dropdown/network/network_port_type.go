package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*NetworkPortType]()
}

// NetworkPortType 网络端口类型
type NetworkPortType struct {
	pkgmodel.Base

	Name *string `json:"name"` // 名称

	// 组织相关
	EntityID    *string `json:"entity_id"`    // 所属实体ID
	IsRecursive *bool   `json:"is_recursive"` // 是否递归继承

	// 配置
	ValueDecimal      *float64 `json:"value_decimal"`      // 十进制值
	IsImportable      *bool    `json:"is_importable"`      // 是否可导入
	InstantiationType *string  `json:"instantiation_type"` // 实例化类型
}

func (*NetworkPortType) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*NetworkPortType]())
}
