package model

import (
	"encoding/json"
	"time"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*PlanningExternalEvent]()
}

// PlanningExternalEvent 外部计划事件
type PlanningExternalEvent struct {
	pkgmodel.Base

	// 基础信息
	UUID         *string        `json:"uuid"`          // 唯一标识符
	Name         *string        `json:"name"`          // 事件名称
	Text         *string        `json:"text"`          // 事件描述
	Status       *PlaningStatus `json:"status"`        // 事件状态
	IsBackground *bool          `json:"is_background"` // 是否后台事件
	IsRecursive  *bool          `json:"is_recursive"`  // 是否继承父级权限

	// 时间信息
	EventDate  *time.Time `json:"event_date"`  // 事件创建日期
	BeginTime  *time.Time `json:"begin_time"`  // 开始时间
	EndTime    *time.Time `json:"end_time"`    // 结束时间
	RepeatRule *string    `json:"repeat_rule"` // 重复规则

	// 关联信息
	TemplateID *string `json:"template_id"` // 模板ID
	CategoryID *string `json:"category_id"` // 事件类别ID
	EntityID   *string `json:"entity_id"`   // 所属实体ID

	// 参与者信息
	UserID   *string          `json:"user_id"`   // 创建用户ID
	GuestIDs *json.RawMessage `json:"guest_ids"` // 访客用户ID列表
	GroupID  *string          `json:"group_id"`  // 关联组ID
}
