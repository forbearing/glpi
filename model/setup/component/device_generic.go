package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*DeviceGeneric]()
}

// DeviceGeneric 通用设备
type DeviceGeneric struct {
	pkgmodel.Base

	Designation *string `json:"designation"` // 设备标识

	// 组织相关
	EntitiesID  *string `json:"entities_id"`  // 所属实体ID
	IsRecursive *bool   `json:"is_recursive"` // 是否递归继承
	LocationID  *string `json:"locations_id"` // 位置ID

	// 制造商信息
	ManufacturerID *string `json:"manufacturers_id"` // 制造商ID
	ModelID        *string `json:"model_id"`         // 型号ID
	TypeID         *string `json:"type_id"`          // 类型ID

	// 状态信息
	StatusID *int `json:"status"` // 状态ID
}

func (*DeviceGeneric) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_COMPONENT, pkgmodel.GetTableName[*DeviceGeneric]())
}
