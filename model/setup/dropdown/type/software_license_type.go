package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*SoftwareLicenseType]()
}

// SoftwareLicenseType 软件许可证类型
type SoftwareLicenseType struct {
	pkgmodel.Base

	// 基础信息
	Name         *string `json:"name"`          // 类型名称
	CompleteName *string `json:"complete_name"` // 完整名称

	// 组织信息
	EntityID    *string `json:"entity_id"`    // 实体ID
	IsRecursive *bool   `json:"is_recursive"` // 是否递归

	// 层级关系
	Level         *int    `json:"level"`          // 层级深度
	ParentID      *int    `json:"parent_id"`      // 父类型ID
	AncestorCache *string `json:"ancestor_cache"` // 祖先缓存
	SonCache      *string `json:"son_cache"`      // 子类缓存
}

func (*SoftwareLicenseType) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*SoftwareLicenseType]())
}
