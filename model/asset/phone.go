package model

import (
	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*Phone]()
}

// Phone 电话资产
type Phone struct {
	pkgmodel.Base

	// 基础信息
	Name            *string           `json:"name"`             // 电话名称
	Status          model.AssetStatus `json:"status"`           // 资产状态
	Brand           *string           `json:"brand"`            // 品牌
	SerialNumber    *string           `json:"serial_number"`    // 序列号
	InventoryNumber *string           `json:"inventory_number"` // 资产编号
	UUID            *string           `json:"uuid"`             // 唯一标识符

	// 组织架构信息
	EntityID    *string `json:"entity_id"`    // 所属实体ID
	LocationID  *string `json:"location_id"`  // 物理位置ID
	IsRecursive *bool   `json:"is_recursive"` // 是否继承父级权限
	IsGlobal    *bool   `json:"is_global"`    // 是否全局可见

	// 配置信息
	ManufacturerID      *string `json:"manufacturer_id"`       // 制造商ID
	PhoneTypeID         *string `json:"phone_type_id"`         // 电话类型ID
	PhoneModelID        *string `json:"phone_model_id"`        // 电话型号ID
	PhonePowerSupplyID  *string `json:"phone_power_supply_id"` // 电源类型ID
	AutoUpdateSystemID  *string `json:"auto_update_system_id"` // 自动更新系统ID
	LastInventoryUpdate *string `json:"last_inventory_update"` // 最后清单更新时间

	// 使用信息
	UserID           *string `json:"user_id"`            // 使用者ID
	GroupID          *string `json:"group_id"`           // 使用组ID
	UserTechnicalID  *string `json:"user_technical_id"`  // 负责运维人员ID
	GroupTechnicalID *string `json:"group_technical_id"` // 负责运维团队ID

	// 联系信息
	ContactName   *string `json:"contact_name"`   // 联系人
	ContactNumber *string `json:"contact_number"` // 联系号码
	NumberLine    *string `json:"number_line"`    // 电话线路号码

	// 功能配置
	HasHeadset *bool `json:"has_headset"` // 是否有耳机
	HasSpeaker *bool `json:"has_speaker"` // 是否有扬声器

	// 模板信息
	IsTemplate   *bool   `json:"is_template"`   // 是否为模板
	TemplateName *string `json:"template_name"` // 模板名称
	IsDynamic    *bool   `json:"is_dynamic"`    // 是否动态

	// 财务信息
	TicketTCO *float64 `json:"ticket_tco"` // 总拥有成本
}
