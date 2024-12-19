package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*DeviceMemory]()
}

// DeviceMemory 内存设备
type DeviceMemory struct {
	pkgmodel.Base

	Designation *string `json:"designation"` // 内存标识

	// 组织相关
	EntityID    *int  `json:"entity_id"`    // 所属实体ID
	IsRecursive *bool `json:"is_recursive"` // 是否递归继承

	// 制造商信息
	ManufacturerID *string `json:"manufacturer_id"` // 制造商ID
	ModelID        *string `json:"model_id"`        // 型号ID

	// 规格参数
	TypeID      *string `json:"type_id"`      // 内存类型ID
	Frequency   *int64  `json:"frequence"`    // 频率(MHz)
	SizeDefault *int64  `json:"size_default"` // 默认容量(MB)
}

func (*DeviceMemory) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_COMPONENT, pkgmodel.GetTableName[*DeviceMemory]())
}
