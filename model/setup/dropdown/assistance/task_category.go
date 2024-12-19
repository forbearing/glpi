package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

// TaskCategory 任务分类信息
type TaskCategory struct {
	pkgmodel.Base

	// 基础信息
	Name         *string `json:"name"`          // 分类名称
	CompleteName *string `json:"complete_name"` // 完整分类名称
	Level        *int    `json:"level"`         // 层级

	// 组织信息
	EntityID    *string `json:"entity_id"`    // 实体ID
	IsRecursive *bool   `json:"is_recursive"` // 是否递归

	// 分类关联
	ParentID           *int `json:"parent_id"`            // 父级分类ID
	KnowbaseCategoryID *int `json:"knowbase_category_id"` // 知识库分类ID

	// 状态标识
	IsActive          *bool `json:"is_active"`           // 是否激活
	IsHelpdeskVisible *bool `json:"is_helpdesk_visible"` // 是否服务台可见

	// 缓存数据
	AncestorCache *string `json:"ancestor_cache"` // 祖先节点缓存
	SonCache      *string `json:"son_cache"`      // 子节点缓存
}

func (*TaskCategory) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*TaskCategory]())
}
