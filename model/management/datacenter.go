package model

import pkgmodel "github.com/forbearing/golib/model"

func init() {
	pkgmodel.Register[*Datacenter]()
}

// Datacenter 数据中心信息
type Datacenter struct {
	pkgmodel.Base

	// 基础信息
	Name        *string `json:"name"`         // 数据中心名称
	IsRecursive *bool   `json:"is_recursive"` // 是否继承父级权限
	Picture     *string `json:"picture"`      // 图片

	// 关联信息
	EntityID   *string `json:"entity_id"`   // 所属实体ID
	LocationID *string `json:"location_id"` // 位置ID
}
