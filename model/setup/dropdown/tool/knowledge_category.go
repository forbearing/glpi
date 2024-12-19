package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*KnowbaseCategory]()
}

// KnowbaseCategory 知识库分类
type KnowbaseCategory struct {
	pkgmodel.Base

	Name         *string `json:"name"`          // 名称
	Level        *int    `json:"level"`         // 层级
	ParentID     *int    `json:"parent_id"`     // 父级ID
	CompleteName *string `json:"complete_name"` // 完整路径名称

	// 组织相关
	EntityID    *string `json:"entity_id"`    // 所属实体ID
	IsRecursive *bool   `json:"is_recursive"` // 是否递归继承

	// 缓存字段
	AncestorCache *string `json:"ancestor_cache"` // 祖先缓存
	SonCache      *string `json:"son_cache"`      // 子级缓存
}

func (*KnowbaseCategory) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*KnowbaseCategory]())
}
