package model

import (
	"fmt"
	"time"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*DeviceFirmware]()
}

// DeviceFirmware 设备固件
type DeviceFirmware struct {
	pkgmodel.Base

	Designation *string `json:"designation"` // 固件标识

	// 组织相关
	EntityID    *string `json:"entity_id"`    // 所属实体ID
	IsRecursive *bool   `json:"is_recursive"` // 是否递归继承

	// 制造商信息
	ManufacturerID *string `json:"manufacturer_id"` // 制造商ID
	ModelID        *string `json:"model_id"`        // 型号ID
	TypeID         *string `json:"type_id"`         // 类型ID

	// 版本信息
	Version     *string    `json:"version"`      // 版本号
	ReleaseDate *time.Time `json:"release_date"` // 发布日期
}

func (*DeviceFirmware) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_COMPONENT, pkgmodel.GetTableName[*DeviceFirmware]())
}
