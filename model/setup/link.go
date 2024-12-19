package model

import pkgmodel "github.com/forbearing/golib/model"

func init() {
	pkgmodel.Register[*Link]()
}

// Link 外部链接
type Link struct {
	pkgmodel.Base

	EntityID    *string `json:"entity_id"`    // 实体ID
	IsRecursive *bool   `json:"is_recursive"` // 是否递归
	Name        *string `json:"name"`         // 链接名称
	Link        *string `json:"link"`         // 链接地址
	Data        *string `json:"data"`         // 附加数据
	OpenWindow  *bool   `json:"open_window"`  // 是否新窗口打开
}
