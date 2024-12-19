package model

import (
	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*Rack]()
}

// Rack 机柜资产
type Rack struct {
	pkgmodel.Base

	// 基础信息
	Name            *string           `json:"name"`             // 机柜名称
	SerialNumber    *string           `json:"serial_number"`    // 序列号
	InventoryNumber *string           `json:"inventory_number"` // 资产编号
	Status          model.AssetStatus `json:"status"`           // 资产状态

	// 组织架构信息
	EntityID    *string `json:"entity_id"`    // 所属实体ID
	LocationID  *string `json:"location_id"`  // 物理位置ID
	DCRoomID    *string `json:"dc_room_id"`   // 数据中心机房ID
	IsRecursive *bool   `json:"is_recursive"` // 是否继承父级权限

	// 配置信息
	ManufacturerID *string `json:"manufacturer_id"` // 制造商ID
	RackModelID    *string `json:"rack_model_id"`   // 机柜型号ID
	RackTypeID     *string `json:"rack_type_id"`    // 机柜类型ID

	// 技术支持信息
	UserTechnicalID  *string `json:"user_technical_id"`  // 负责运维人员ID
	GroupTechnicalID *string `json:"group_technical_id"` // 负责运维团队ID

	// 物理参数
	Width         *int64   `json:"width"`          // 宽度(mm)
	Height        *int64   `json:"height"`         // 高度(mm)
	Depth         *int64   `json:"depth"`          // 深度(mm)
	NumberUnit    *int64   `json:"number_unit"`    // 机柜单元数
	Position      *string  `json:"position"`       // 位置信息
	Orientation   *int64   `json:"orientation"`    // 方向
	BGColor       *string  `json:"bg_color"`       // 背景颜色
	MaxPower      *float64 `json:"max_power"`      // 最大功率(W)
	MeasuredPower *float64 `json:"measured_power"` // 实测功率(W)
	MaxWeight     *float64 `json:"max_weight"`     // 最大承重(kg)

	// 模板信息
	IsTemplate   *bool   `json:"is_template"`   // 是否为模板
	TemplateName *string `json:"template_name"` // 模板名称
}
