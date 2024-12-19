package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*DeviceCase]()
}

// DeviceCase 设备机箱
type DeviceCase struct {
	pkgmodel.Base

	Designation *string `json:"designation"` // 机箱标识

	// 组织相关
	EntityID    *string `json:"entity_id"`    // 所属实体ID
	IsRecursive *bool   `json:"is_recursive"` // 是否递归继承

	// 制造商信息
	ManufacturerID *string `json:"manufacturer_id"` // 制造商ID
	ModelID        *string `json:"model_id"`        // 型号ID
	TypeID         *string `json:"type_id"`         // 类型ID
}

func (*DeviceCase) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_COMPONENT, pkgmodel.GetTableName[*DeviceCase]())
}
