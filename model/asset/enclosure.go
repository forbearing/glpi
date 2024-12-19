package model

import (
	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*Enclosure]()
}

// Enclosure 机箱资产
type Enclosure struct {
	pkgmodel.Base

	// 基础信息
	Name            *string           `json:"name"`             // 机箱名称
	SerialNumber    *string           `json:"serial_number"`    // 序列号
	InventoryNumber *string           `json:"inventory_number"` // 资产编号
	Status          model.AssetStatus `json:"status"`           // 资产状态

	// 组织架构信息
	EntityID    *string `json:"entity_id"`    // 所属实体ID
	LocationID  *string `json:"location_id"`  // 物理位置ID
	IsRecursive *bool   `json:"is_recursive"` // 是否继承父级权限

	// 配置信息
	ManufacturerID   *string `json:"manufacturer_id"`    // 制造商ID
	EnclosureModelID *string `json:"enclosure_model_id"` // 机箱型号ID

	// 技术支持信息
	UserTechnicalID  *string `json:"user_technical_id"`  // 负责运维人员ID
	GroupTechnicalID *string `json:"group_technical_id"` // 负责运维团队ID

	// 物理参数
	Orientation   *int64 `json:"orientation"`    // 方向
	PowerSupplies *int64 `json:"power_supplies"` // 电源数量

	// 模板信息
	IsTemplate   *bool   `json:"is_template"`   // 是否为模板
	TemplateName *string `json:"template_name"` // 模板名称
}
