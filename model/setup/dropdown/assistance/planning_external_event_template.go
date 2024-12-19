package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*PlanningExternalEventTemplate]()
}

// PlanningExternalEventTemplate 外部事件模板
type PlanningExternalEventTemplate struct {
	pkgmodel.Base

	// 基础信息
	Name   *string                            `json:"name"`   // 模板名称
	Text   *string                            `json:"text"`   // 事件内容
	Status *PlanningExternalEventTemplateType `json:"status"` // 状态

	// 组织信息
	EntityID *string `json:"entity_id"` // 实体ID

	// 时间信息
	Duration   *int    `json:"duration"`    // 持续时间
	BeforeTime *int    `json:"before_time"` // 提前时间
	RepeatRule *string `json:"repeat_rule"` // 重复规则

	// 状态和分类
	Background *bool `json:"background"`  // 是否后台
	CategoryID *int  `json:"category_id"` // 事件分类ID
}

func (*PlanningExternalEventTemplate) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*PlanningExternalEventTemplate]())
}
