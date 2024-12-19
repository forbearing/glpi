package model

import (
	"time"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*Project]()
}

// Project 项目信息
type Project struct {
	pkgmodel.Base

	// 基础信息
	Name        *string          `json:"name"`         // 项目名称
	Code        *string          `json:"code"`         // 项目代码
	Priority    *ProjectPriority `json:"priority"`     // 优先级
	Status      *ProjectStatus   `json:"status"`       // 项目状态ID
	Content     *string          `json:"content"`      // 项目内容
	PercentDone *int             `json:"percent_done"` // 完成百分比

	// 计划时间
	PlanStartDate *time.Time `json:"plan_start_date"` // 计划开始时间
	PlanEndDate   *time.Time `json:"plan_end_date"`   // 计划结束时间

	// 实际时间
	RealStartDate *time.Time `json:"real_start_date"` // 实际开始时间
	RealEndDate   *time.Time `json:"real_end_date"`   // 实际结束时间

	// 配置选项
	IsRecursive       *bool `json:"is_recursive"`         // 是否继承父级权限
	AutoPercentDone   *bool `json:"auto_percent_done"`    // 是否自动计算完成度
	ShowOnGlobalGantt *bool `json:"show_on_global_gantt"` // 是否显示在全局甘特图

	// 模板信息
	IsTemplate   *bool   `json:"is_template"`   // 是否为模板
	TemplateName *string `json:"template_name"` // 模板名称

	// 关联信息
	EntityID          *string `json:"entity_id"`           // 所属实体ID
	ParentProjectID   *string `json:"parent_project_id"`   // 父项目ID
	ProjectTypeID     *string `json:"project_type_id"`     // 项目类型ID
	ProjectTemplateID *string `json:"project_template_id"` // 项目模板ID
	UserID            *string `json:"user_id"`             // 用户ID
	GroupID           *string `json:"group_id"`            // 组ID
}
