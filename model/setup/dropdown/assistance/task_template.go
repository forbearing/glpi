package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*TaskTemplate]()
}

// TaskTemplate 任务模板信息
type TaskTemplate struct {
	pkgmodel.Base

	// 基础信息
	Name       *string           `json:"name"`        // 模板名称
	Content    *string           `json:"content"`     // 模板内容
	Status     *TaskTemplateType `json:"status"`      // 状态
	ActionTime *int              `json:"action_time"` // 操作时间,单位:秒

	// 组织信息
	EntityID    *string `json:"entity_id"`    // 实体ID
	IsRecursive *bool   `json:"is_recursive"` // 是否递归
	IsPrivate   *bool   `json:"is_private"`   // 是否私有

	// 分类关联
	ParentID *int `json:"parent_id"` // 父类ID

	// 技术支持
	TechUserID  *int `json:"tech_user_id"`  // 技术支持用户ID
	TechGroupID *int `json:"tech_group_id"` // 技术支持组ID
}

func (*TaskTemplate) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*TaskTemplate]())
}
