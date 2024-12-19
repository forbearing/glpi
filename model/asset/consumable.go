package model

import (
	"time"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*Consumable]()
	pkgmodel.Register[*ConsumableItem]()
}

// Consumable 耗材模板
type Consumable struct {
	pkgmodel.Base

	// 基础信息
	Name            *string `json:"name"`             // 耗材名称
	Reference       *string `json:"reference"`        // 耗材参考编号
	InventoryNumber *string `json:"inventory_number"` // 资产编号
	Picture         *string `json:"picture"`          // 耗材图片

	// 组织架构信息
	EntityID    *string `json:"entity_id"`    // 所属实体ID
	LocationID  *string `json:"location_id"`  // 物理位置ID
	IsRecursive *bool   `json:"is_recursive"` // 是否继承父级权限

	// 配置信息
	ManufacturerID   *string `json:"manufacturer_id"`    // 制造商ID
	ConsumableTypeID *string `json:"consumable_type_id"` // 耗材类型ID

	// 库存信息
	AlarmThreshold *int64 `json:"alarm_threshold"` // 库存警告阈值
	StockTarget    *int64 `json:"stock_target"`    // 目标库存量

	// 技术支持信息
	UserTechnicalID  *string `json:"user_technical_id"`  // 负责运维人员ID
	GroupTechnicalID *string `json:"group_technical_id"` // 负责运维团队ID
}

// ConsumableItem 耗材实例
type ConsumableItem struct {
	pkgmodel.Base

	// 基础信息
	ConsumableID *string `json:"consumable_id"` // 关联的耗材模板ID
	EntityID     *string `json:"entity_id"`     // 所属实体ID

	// 使用信息
	DateIn   *time.Time `json:"date_in"`   // 入库日期
	DateOut  *time.Time `json:"date_out"`  // 出库日期
	ItemType *string    `json:"item_type"` // 关联设备类型
	ItemID   *string    `json:"item_id"`   // 关联设备ID
}
