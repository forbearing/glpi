package model

import (
	"time"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*Knowbase]()
}

// Knowbase 知识库条目
type Knowbase struct {
	pkgmodel.Base

	// 基础信息
	Name    *string `json:"name"`    // 条目名称
	Content *string `json:"content"` // 条目内容
	IsFAQ   *bool   `json:"is_faq"`  // 是否常见问题
	Views   *int    `json:"views"`   // 浏览次数

	// 时间范围
	VisibleBeginDate *time.Time `json:"visible_begin_date"` // 可见开始时间
	VisibleEndDate   *time.Time `json:"visible_end_date"`   // 可见结束时间

	// 关联信息
	UserID *string `json:"user_id"` // 用户ID
}
