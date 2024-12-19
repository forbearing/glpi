package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*DeviceSensor]()
}

// DeviceSensor 传感器设备
type DeviceSensor struct {
	pkgmodel.Base

	Designation *string `json:"designation"` // 传感器标识

	// 组织相关
	EntityID    *string `json:"entity_id"`    // 所属实体ID
	IsRecursive *bool   `json:"is_recursive"` // 是否递归继承

	// 位置状态
	LocationID *string `json:"location_id"` // 位置ID
	StatusID   *string `json:"status_id"`   // 状态ID

	// 制造商信息
	ManufacturerID *string `json:"manufacturer_id"` // 制造商ID
	ModelID        *string `json:"model_id"`        // 型号ID
	TypeID         *string `json:"type_id"`         // 传感器类型ID
}

func (*DeviceSensor) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_COMPONENT, pkgmodel.GetTableName[*DeviceSensor]())
}
