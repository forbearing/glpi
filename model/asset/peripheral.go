package model

import (
	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*Peripheral]()
}

type Peripheral struct {
	pkgmodel.Base

	// 基础信息
	Name            *string           `json:"name"`             // 外设名称
	UUID            *string           `json:"uuid"`             // 设备唯一标识符
	Status          model.AssetStatus `json:"status"`           // 资产状态
	SerialNumber    *string           `json:"serial_number"`    // 硬件序列号
	InventoryNumber *string           `json:"inventory_number"` // 资产编号
	Brand           *string           `json:"brand"`            // 品牌名称

	// 组织架构信息
	EntityID    *string `json:"entity_id"`    // 所属实体ID
	LocationID  *string `json:"location_id"`  // 物理位置ID
	IsRecursive *bool   `json:"is_recursive"` // 是否继承父级权限
	IsGlobal    *bool   `json:"is_global"`    // 是否为全局设备

	// 配置信息
	ManufacturerID    *string `json:"manufacturer_id"`     // 制造商ID
	PeripheralTypeID  *string `json:"peripheral_type_id"`  // 外设类型ID
	PeripheralModelID *string `json:"peripheral_model_id"` // 外设型号ID

	// 使用者信息
	UserID        *string `json:"user_id"`        // 主要使用者ID
	GroupID       *string `json:"group_id"`       // 使用者所属组ID
	ContactName   *string `json:"contact_name"`   // 联系人姓名
	ContactNumber *string `json:"contact_number"` // 联系人工号/电话

	// 技术支持信息
	UserTechnicalID  *string `json:"user_technical_id"`  // 负责运维人员ID
	GroupTechnicalID *string `json:"group_technical_id"` // 负责运维团队ID

	// 模版信息
	IsTemplate   *bool   `json:"is_template"`   // 是否为模板记录
	TemplateName *string `json:"template_name"` // 模板名称

	// 资产同步信息
	IsDynamic          *bool   `json:"is_dynamic"`            // 是否自动更新
	AutoUpdateSystemID *string `json:"auto_update_system_id"` // 更新来源ID

	// 财务信息
	TicketTCO *float64 `json:"ticket_tco"` // 总拥有成本(TCO)
}
