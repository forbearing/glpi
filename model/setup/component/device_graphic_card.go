package model

import pkgmodel "github.com/forbearing/golib/model"

// DeviceGraphicCard 显卡设备
type DeviceGraphicCard struct {
	pkgmodel.Base

	Designation *string `json:"designation"` // 显卡标识

	// 组织相关
	EntityID    *int  `json:"entity_id"`    // 所属实体ID
	IsRecursive *bool `json:"is_recursive"` // 是否递归继承

	// 制造商信息
	ManufacturerID *int `json:"manufacturer_id"` // 制造商ID
	ModelID        *int `json:"model_id"`        // 型号ID

	// 硬件规格
	InterfaceTypeID *int    `json:"interface_type_id"` // 接口类型ID
	Chipset         *string `json:"chipset"`           // 芯片组
	MemoryDefault   *int    `json:"memory_default"`    // 默认显存大小(MB)
}
