package model

import (
	"time"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*ChangeRecurrent]()
}

// ChangeRecurrent 周期性变更
type ChangeRecurrent struct {
	pkgmodel.Base

	// 基础信息
	Name             *string `json:"name"`               // 名称
	EntityID         *string `json:"entity_id"`          // 实体ID
	IsRecursive      *bool   `json:"is_recursive"`       // 是否递归
	IsActive         *bool   `json:"is_active"`          // 是否激活
	ChangeTemplateID *int    `json:"change_template_id"` // 变更模板ID

	// 周期设置
	BeginDate        *time.Time `json:"begin_date"`         // 开始日期
	EndDate          *time.Time `json:"end_date"`           // 结束日期
	Periodicity      *int       `json:"periodicity"`        // 周期(秒)
	BeforeCreate     *int       `json:"before_create"`      // 提前创建时间
	NextCreationDate *time.Time `json:"next_creation_date"` // 下次创建日期
	CalendarID       *int       `json:"calendar_id"`        // 日历ID
}
