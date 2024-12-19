package model

import (
	"time"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*TicketRecurrent]()
}

// TicketRecurrent 周期性工单
type TicketRecurrent struct {
	pkgmodel.Base
	// 基础信息
	Name             *string `json:"name"`               // 名称
	EntityD          *string `json:"entity_id"`          // 实体ID
	IsRecursive      *bool   `json:"is_recursive"`       // 是否递归
	IsActive         *bool   `json:"is_active"`          // 是否激活
	TicketTemplateID *int    `json:"ticket_template_id"` // 工单模板ID

	// 周期设置
	BeginDate        *time.Time `json:"begin_date"`         // 开始日期
	EndDate          *time.Time `json:"end_date"`           // 结束日期
	Periodicity      *int       `json:"periodicity"`        // 周期(秒)
	BeforeCreate     *int       `json:"before_create"`      // 提前创建时间
	NextCreationDate *time.Time `json:"next_creation_date"` // 下次创建日期
	CalendarID       *int       `json:"calendar_id"`        // 日历ID
}
