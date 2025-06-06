package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*DeviceBattery]()
}

// DeviceBattery 设备电池
type DeviceBattery struct {
	pkgmodel.Base

	Designation *string `json:"designation"` // 电池标识

	// 组织相关
	EntityID    *string `json:"entity_id"`    // 所属实体ID
	IsRecursive *bool   `json:"is_recursive"` // 是否递归继承

	// 制造商信息
	ManufacturerID *string `json:"manufacturer_id"` // 制造商ID
	ModelID        *string `json:"model_id"`        // 型号ID
	TypeID         *string `json:"type_id"`         // 类型ID

	// 电池参数
	Voltage  *float64 `json:"voltage"`  // 电压(V)
	Capacity *int64   `json:"capacity"` // 容量(mAh)
}

func (*DeviceBattery) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_COMPONENT, pkgmodel.GetTableName[*DeviceBattery]())
}
