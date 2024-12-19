package model

import (
	"time"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*Ticket]()
}

// Ticket 工单信息
type Ticket struct {
	pkgmodel.Base

	// 基础信息
	Name    *string       `json:"name"`    // 工单标题
	Content *string       `json:"content"` // 工单内容
	Type    *TicketType   `json:"type"`    // 工单类型
	Status  *TicketStatus `json:"status"`  // 工单状态

	// 时间信息
	OpenDate            *time.Time `json:"open_date"`              // 创建时间
	CloseDate           *time.Time `json:"close_date"`             // 关闭时间
	SolveDate           *time.Time `json:"solve_date"`             // 解决时间
	TakeIntoAccountDate *time.Time `json:"take_into_account_date"` // 接单时间
	BeginWaitingDate    *time.Time `json:"begin_waiting_date"`     // 等待开始时间
	TimeToResolve       *time.Time `json:"time_to_resolve"`        // 解决期限
	TimeToOwn           *time.Time `json:"time_to_own"`            // 响应期限

	// 优先级信息
	Urgency  *Urgency  `json:"urgency"`  // 紧急程度
	Impact   *Impact   `json:"impact"`   // 影响程度
	Priority *Priority `json:"priority"` // 优先级别

	// 分类信息
	CategoryID    *string `json:"category_id"`     // 工单类别ID
	RequestTypeID *string `json:"request_type_id"` // 请求类型ID
	LocationID    *string `json:"location_id"`     // 位置ID

	// 用户信息
	LastUpdaterID *string `json:"last_updater_id"` // 最后更新人ID
	RecipientID   *string `json:"recipient_id"`    // 接收人ID
	EntityID      *string `json:"entity_id"`       // 所属实体ID

	// SLA相关
	SLATtrID           *string    `json:"sla_ttr_id"`           // SLA解决时间ID
	SLATtoID           *string    `json:"sla_tto_id"`           // SLA响应时间ID
	SLALevelTtrID      *string    `json:"sla_level_ttr_id"`     // SLA解决等级ID
	SLAWaitingDuration *int64     `json:"sla_waiting_duration"` // SLA等待时长
	OLAWaitingDuration *int64     `json:"ola_waiting_duration"` // OLA等待时长
	OLATtoID           *string    `json:"ola_tto_id"`           // OLA响应时间ID
	OLATtrID           *string    `json:"ola_ttr_id"`           // OLA解决时间ID
	OLALevelTtrID      *string    `json:"ola_level_ttr_id"`     // OLA解决等级ID
	OLATtoBeginDate    *time.Time `json:"ola_tto_begin_date"`   // OLA响应开始时间
	OLATtrBeginDate    *time.Time `json:"ola_ttr_begin_date"`   // OLA解决开始时间

	// 统计信息
	ValidationPercent        *int   `json:"validation_percent"`           // 验证百分比
	WaitingDuration          *int64 `json:"waiting_duration"`             // 等待时长
	ActionTime               *int64 `json:"action_time"`                  // 操作时长
	CloseDelayStat           *int64 `json:"close_delay_stat"`             // 关闭延迟统计
	SolveDelayStat           *int64 `json:"solve_delay_stat"`             // 解决延迟统计
	TakeIntoAccountDelayStat *int64 `json:"take_into_account_delay_stat"` // 接单延迟统计
}
