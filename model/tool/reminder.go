package model

import (
	"time"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*Reminder]()
}

// Reminder 提醒信息
type Reminder struct {
	pkgmodel.Base

	// 基础信息
	Name      *string         `json:"name"`       // 提醒名称
	UUID      *string         `json:"uuid"`       // 唯一标识符
	Text      *string         `json:"text"`       // 提醒内容
	Status    *ReminderStatus `json:"status"`     // 状态
	IsPlanned *bool           `json:"is_planned"` // 是否已计划

	// 时间范围
	BeginTime *time.Time `json:"begin_time"` // 开始时间
	EndTime   *time.Time `json:"end_time"`   // 结束时间

	// 显示范围
	ViewBeginTime *time.Time `json:"view_begin_time"` // 显示开始时间
	ViewEndTime   *time.Time `json:"view_end_time"`   // 显示结束时间

	// 关联信息
	UserID *string `json:"user_id"` // 用户ID
}
