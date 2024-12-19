package model

import (
	"time"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*SavedSearch]()
}

// SavedSearch 保存的搜索信息
type SavedSearch struct {
	pkgmodel.Base

	// 基础信息
	Name      *string `json:"name"`       // 搜索名称
	Type      *int    `json:"type"`       // 搜索类型
	ItemType  *string `json:"item_type"`  // 搜索对象类型
	IsPrivate *bool   `json:"is_private"` // 是否私有
	Query     *string `json:"query"`      // 搜索查询条件

	// 组织信息
	UserID      *int    `json:"user_id"`      // 用户ID
	EntityID    *string `json:"entity_id"`    // 实体ID
	IsRecursive *bool   `json:"is_recursive"` // 是否递归

	// 执行统计
	Counter           *int       `json:"counter"`             // 执行次数
	DoCount           *int       `json:"do_count"`            // 计数方式
	LastExecutionTime *float64   `json:"last_execution_time"` // 最后执行时长
	LastExecutionDate *time.Time `json:"last_execution_date"` // 最后执行时间
}
