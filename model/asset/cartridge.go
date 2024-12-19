package model

import pkgmodel "github.com/forbearing/golib/model"

func init() {
	pkgmodel.Register[*Cartridge]()
}

type Cartridge struct {
	pkgmodel.Base

	// 基础信息
	Name      *string `json:"name"`      // 耗材名称
	Reference *string `json:"reference"` // 耗材参考编号
	Picture   *string `json:"picture"`   // 耗材图片

	// 组织架构信息
	EntityID    *string `json:"entity_id"`    // 所属实体ID
	LocationID  *string `json:"location_id"`  // 物理位置ID
	IsRecursive *bool   `json:"is_recursive"` // 是否继承父级权限

	// 配置信息
	ManufacturerID      *string `json:"manufacturer_id"`        // 制造商ID
	CartridgeItemTypeID *string `json:"cartridge_item_type_id"` // 耗材类型ID

	// 库存信息
	AlarmThreshold *int64 `json:"alarm_threshold"` // 库存警告阈值
	StockTarget    *int64 `json:"stock_target"`    // 目标库存量

	// 技术支持信息
	UserTechnicalID  *string `json:"user_technical_id"`  // 负责运维人员ID
	GroupTechnicalID *string `json:"group_technical_id"` // 负责运维团队ID
}
