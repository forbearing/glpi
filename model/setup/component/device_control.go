package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*DeviceControl]()
}

// DeviceControl 设备控制器
type DeviceControl struct {
	pkgmodel.Base

	Designation *string `json:"designation"` // 控制器标识

	// 组织相关
	EntityID    *string `json:"entity_id"`    // 所属实体ID
	IsRecursive *bool   `json:"is_recursive"` // 是否递归继承

	// 制造商信息
	ManufacturerID *string `json:"manufacturer_id"` // 制造商ID
	ModelID        *string `json:"model_id"`        // 型号ID

	// 接口信息
	InterfaceTypeID *string `json:"interface_type_id"` // 接口类型ID

	// 功能特性
	IsRAID *bool `json:"is_raid"` // 是否支持RAID
}

func (*DeviceControl) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_COMPONENT, pkgmodel.GetTableName[*DeviceControl]())
}
