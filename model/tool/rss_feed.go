package model

import pkgmodel "github.com/forbearing/golib/model"

func init() {
	pkgmodel.Register[*RSSFeed]()
}

// RSSFeed RSS订阅源
type RSSFeed struct {
	pkgmodel.Base

	// 基础信息
	Name      *string `json:"name"`       // 订阅源名称
	URL       *string `json:"url"`        // 订阅源地址
	IsActive  *bool   `json:"is_active"`  // 是否激活
	HaveError *bool   `json:"have_error"` // 是否有错误

	// 配置参数
	RefreshRate *int64 `json:"refresh_rate"` // 刷新间隔(秒)
	MaxItems    *int   `json:"max_items"`    // 最大条目数

	// 关联信息
	UserID *string `json:"user_id"` // 用户ID
}
