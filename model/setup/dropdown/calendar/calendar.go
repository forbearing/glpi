package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*Calendar]()
}

// Calendar 日历
type Calendar struct {
	pkgmodel.Base

	Name *string `json:"name"` // 名称

	// 组织相关
	EntityID    *string `json:"entity_id"`    // 所属实体ID
	IsRecursive *bool   `json:"is_recursive"` // 是否递归继承

	// 缓存设置
	CacheDuration *string `json:"cache_duration"` // 每周各天的缓存时长
}

func (*Calendar) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*Calendar]())
}
