package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*PendingReason]()
}

// PendingReason 待处理原因
type PendingReason struct {
	pkgmodel.Base

	// 基础信息
	Name *string `json:"name"` // 原因名称

	// 组织信息
	EntityID    *string `json:"entity_id"`    // 实体ID
	IsRecursive *bool   `json:"is_recursive"` // 是否递归

	// 跟进配置
	FollowupFrequency        *int `json:"followup_frequency"`         // 跟进频率
	FollowupBeforeResolution *int `json:"followup_before_resolution"` // 解决前跟进次数

	// 关联模板
	ITILFollowupTemplateID *int `json:"itil_followup_template_id"` // 跟进模板ID
	SolutionTemplateID     *int `json:"solution_template_id"`      // 解决方案模板ID
}

func (*PendingReason) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*PendingReason]())
}
