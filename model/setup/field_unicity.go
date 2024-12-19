package model

import pkgmodel "github.com/forbearing/golib/model"

func init() {
	pkgmodel.Register[*FieldUnicity]()
}

// FieldUnicity 字段唯一性
type FieldUnicity struct {
	pkgmodel.Base

	Name *string `json:"name"` // 名称

	// 组织相关
	EntityID    *string `json:"entity_id"`    // 所属实体ID
	IsRecursive *bool   `json:"is_recursive"` // 是否递归继承

	// 字段相关
	ItemType *string `json:"item_type"` // 对象类型
	Fields   *string `json:"fields"`    // 字段名称
	IsActive *bool   `json:"is_active"` // 是否激活

	// 动作相关
	ActionRefuse *bool `json:"action_refuse"` // 是否拒绝重复, Record into the database denied
	ActionNotify *bool `json:"action_notify"` // 是否发送通知
}
