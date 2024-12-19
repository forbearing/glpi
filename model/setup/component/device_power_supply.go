package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*DevicePowerSupply]()
}

// DevicePowerSupply 电源设备
type DevicePowerSupply struct {
	pkgmodel.Base

	Designation *string `json:"designation"` // 电源标识

	// 组织相关
	EntityID    *string `json:"entity_id"`    // 所属实体ID
	IsRecursive *bool   `json:"is_recursive"` // 是否递归继承

	// 制造商信息
	ManufacturerID *string `json:"manufacturer_id"` // 制造商ID
	ModelID        *string `json:"model_id"`        // 型号ID

	// 规格参数
	Power *string `json:"power"`  // 功率
	IsATX *bool   `json:"is_atx"` // 是否ATX规格
}

func (*DevicePowerSupply) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_COMPONENT, pkgmodel.GetTableName[*DevicePowerSupply]())
}
