package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*DeviceMotherboard]()
}

// DeviceMotherboard 主板设备
type DeviceMotherboard struct {
	pkgmodel.Base

	Designation *string `json:"designation"` // 主板标识
	Chipset     *string `json:"chipset"`     // 芯片组

	// 组织相关
	EntityID    *int  `json:"entity_id"`    // 所属实体ID
	IsRecursive *bool `json:"is_recursive"` // 是否递归继承

	// 制造商信息
	ManufacturerID *int `json:"manufacturer_id"` // 制造商ID
	ModelID        *int `json:"model_id"`        // 型号ID
}

func (*DeviceMotherboard) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_COMPONENT, pkgmodel.GetTableName[*DeviceMotherboard]())
}
