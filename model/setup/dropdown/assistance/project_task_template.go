package model

import (
	"fmt"
	"time"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*ProjectTaskTemplate]()
}

// ProjectTaskTemplate 项目任务模板
type ProjectTaskTemplate struct {
	pkgmodel.Base

	// 基础信息
	Name        *string `json:"name"`        // 模板名称
	Description *string `json:"description"` // 任务描述

	// 组织信息
	EntityID    *string `json:"entity_id"`    // 实体ID
	IsRecursive *bool   `json:"is_recursive"` // 是否递归

	// 关联信息
	ProjectID       *int `json:"project_id"`           // 项目ID
	ParentTaskID    *int `json:"project_task_id"`      // 父任务ID
	ProjectStatusID *int `json:"project_status_id"`    // 项目状态ID
	TaskTypeID      *int `json:"project_task_type_id"` // 任务类型ID
	UserID          *int `json:"users_id"`             // 用户ID

	// 时间信息
	PlanStartDate *time.Time `json:"plan_start_date"` // 计划开始时间
	PlanEndDate   *time.Time `json:"plan_end_date"`   // 计划结束时间
	RealStartDate *time.Time `json:"real_start_date"` // 实际开始时间
	RealEndDate   *time.Time `json:"real_end_date"`   // 实际结束时间

	// 进度信息
	PlannedDuration   *int  `json:"planned_duration"`   // 计划持续时间
	EffectiveDuration *int  `json:"effective_duration"` // 实际持续时间
	PercentDone       *int  `json:"percent_done"`       // 完成百分比
	IsMilestone       *bool `json:"is_milestone"`       // 是否里程碑
}

func (*ProjectTaskTemplate) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*ProjectTaskTemplate]())
}
