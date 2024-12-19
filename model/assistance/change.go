package model

import (
	"time"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*Change]()
}

// Change 变更信息
type Change struct {
	pkgmodel.Base

	// 基础信息
	Name        *string       `json:"name"`         // 变更标题
	Content     *string       `json:"content"`      // 变更描述
	Status      *ChangeStatus `json:"status"`       // 变更状态
	IsRecursive *bool         `json:"is_recursive"` // 是否继承父级权限

	// 时间信息
	OpenDate         *time.Time `json:"open_date"`          // 创建时间
	CloseDate        *time.Time `json:"close_date"`         // 关闭时间
	SolveDate        *time.Time `json:"solve_date"`         // 解决时间
	TimeToResolve    *time.Time `json:"time_to_resolve"`    // 解决期限
	BeginWaitingDate *time.Time `json:"begin_waiting_date"` // 等待开始时间

	// 优先级信息
	Urgency  *Urgency  `json:"urgency"`  // 紧急程度
	Impact   *Impact   `json:"impact"`   // 影响程度
	Priority *Priority `json:"priority"` // 优先级别

	// 变更内容
	ImpactContent      *string `json:"impact_content"`       // 影响评估内容
	ControlListContent *string `json:"control_list_content"` // 控制清单内容
	RolloutPlanContent *string `json:"rollout_plan_content"` // 实施计划内容
	BackoutPlanContent *string `json:"backout_plan_content"` // 回退计划内容
	CheckListContent   *string `json:"check_list_content"`   // 检查清单内容

	// 分类信息
	CategoryID *string `json:"category_id"` // 变更类别ID
	LocationID *string `json:"location_id"` // 位置ID
	EntityID   *string `json:"entity_id"`   // 所属实体ID

	// 用户信息
	LastUpdaterID *string `json:"last_updater_id"` // 最后更新人ID
	RecipientID   *string `json:"recipient_id"`    // 接收人ID

	// 验证信息
	GlobalValidation  *int `json:"global_validation"`  // 全局验证状态
	ValidationPercent *int `json:"validation_percent"` // 验证进度百分比

	// 统计信息
	ActionTime      *int64 `json:"action_time"`      // 操作时长
	WaitingDuration *int64 `json:"waiting_duration"` // 等待时长
	CloseDelayStat  *int64 `json:"close_delay_stat"` // 关闭延迟统计
	SolveDelayStat  *int64 `json:"solve_delay_stat"` // 解决延迟统计
}
