package model

import (
	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*Appliance]()
}

// Appliance 应用设备信息
type Appliance struct {
	pkgmodel.Base

	// 基础信息
	Name              *string            `json:"name"`                // 设备名称
	SerialNumber      *string            `json:"serial_number"`       // 序列号
	InventorySerial   *string            `json:"inventory_serial"`    // 其他序列号
	ExternalID        *string            `json:"external_id"`         // 外部标识符
	Status            *model.AssetStatus `json:"status"`              // 状态
	IsHelpDeskVisible *bool              `json:"is_helpdesk_visible"` // 是否在服务台可见
	IsRecursive       *bool              `json:"is_recursive"`        // 是否继承父级权限
	Picture           *string            `json:"picture"`             // 图片

	// 联系信息
	ContactName   *string `json:"contact_name"`   // 联系人
	ContactNumber *string `json:"contact_number"` // 联系电话

	// 关联信息
	EntityID               *string `json:"entity_id"`                // 所属实体ID
	LocationID             *string `json:"location_id"`              // 位置ID
	ManufacturerID         *string `json:"manufacturer_id"`          // 制造商ID
	ApplianceTypeID        *string `json:"appliance_type_id"`        // 设备类型ID
	ApplianceEnvironmentID *string `json:"appliance_environment_id"` // 设备环境ID
	UserID                 *string `json:"user_id"`                  // 用户ID
	GroupID                *string `json:"group_id"`                 // 组ID
	TechUserID             *string `json:"tech_user_id"`             // 技术用户ID
	TechGroupID            *string `json:"tech_group_id"`            // 技术组ID
}
