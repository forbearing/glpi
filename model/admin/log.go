package model

import pkgmodel "github.com/forbearing/golib/model"

func init() {
	pkgmodel.Register[*Log]()
}

// Log 日志信息
type Log struct {
	pkgmodel.Base

	// 关联对象
	ItemType     *string `json:"item_type"`      // 对象类型
	ItemID       *int    `json:"item_id"`        // 对象ID
	ItemTypeLink *string `json:"item_type_link"` // 关联对象类型

	// 操作信息
	LinkedAction *int    `json:"linked_action"` // 关联动作
	UserName     *string `json:"user_name"`     // 操作用户

	// 变更内容
	SearchOptionID *int    `json:"search_option_id"` // 搜索选项ID
	OldValue       *string `json:"old_value"`        // 原值
	NewValue       *string `json:"new_value"`        // 新值
}
