package model

import pkgmodel "github.com/forbearing/golib/model"

// FieldBlacklist 字段黑名单
type FieldBlacklist struct {
	pkgmodel.Base

	Name     *string `json:"name"`      // 黑名单名称
	Field    *string `json:"field"`     // 字段名
	Value    *string `json:"value"`     // 字段值
	ItemType *string `json:"item_type"` // 项目类型

	// 组织相关
	Entity      *string `json:"entity_id"`    // 所属实体ID
	IsRecursive *bool   `json:"is_recursive"` // 是否递归继承
}
