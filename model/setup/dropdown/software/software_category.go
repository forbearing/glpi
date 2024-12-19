package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*SoftwareCategory]()
}

// SoftwareCategory 软件分类
type SoftwareCategory struct {
	pkgmodel.Base

	Name         *string `json:"name"`          // 名称
	CompleteName *string `json:"complete_name"` // 完整名称
	Level        *int    `json:"level"`         // 层级

	// 分类关系
	ParentID      *int    `json:"parent_id"`      // 父分类ID
	AncestorCache *string `json:"ancestor_cache"` // 祖先缓存
	SonCache      *string `json:"son_cache"`      // 子分类缓存
}

func (*SoftwareCategory) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*SoftwareCategory]())
}
