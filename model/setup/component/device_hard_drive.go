package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*DeviceHardDrive]()
}

// DeviceHardDrive 硬盘驱动器
type DeviceHardDrive struct {
	pkgmodel.Base

	Designation *string `json:"designation"` // 硬盘标识

	// 组织相关
	EntityID    *int  `json:"entity_id"`    // 所属实体ID
	IsRecursive *bool `json:"is_recursive"` // 是否递归继承

	// 制造商信息
	ManufacturerID *int `json:"manufacturer_id"` // 制造商ID
	ModelID        *int `json:"model_id"`        // 型号ID

	// 硬件规格
	InterfaceTypeID *int    `json:"interface_type_id"` // 接口类型ID
	RPM             *string `json:"rpm"`               // 转速
	Cache           *int64  `json:"cache"`             // 缓存大小
	CapacityDefault *int64  `json:"capacity_default"`  // 默认容量(GB)
}

func (*DeviceHardDrive) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_COMPONENT, pkgmodel.GetTableName[*DeviceHardDrive]())
}
