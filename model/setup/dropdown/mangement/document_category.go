package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*DocumentCategory]()
}

// DocumentCategory 文档分类
type DocumentCategory struct {
	pkgmodel.Base

	Name         *string `json:"name"`          // 名称
	Level        *int    `json:"level"`         // 层级
	ParentID     *int    `json:"parent_id"`     // 父分类ID
	CompleteName *string `json:"complete_name"` // 完整路径名称

	// 缓存字段
	AncestorCache *string `json:"ancestor_cache"` // 祖先缓存
	SonCache      *string `json:"son_cache"`      // 子分类缓存
}

func (*DocumentCategory) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*DocumentCategory]())
}
