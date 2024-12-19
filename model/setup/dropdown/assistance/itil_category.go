package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*ITILCategory]()
}

// ITILCategory ITIL分类信息
type ITILCategory struct {
	pkgmodel.Base

	// 基础信息
	Name         *string `json:"name"`          // 分类名称
	CompleteName *string `json:"complete_name"` // 完整分类名称
	Level        *int    `json:"level"`         // 层级
	Code         *string `json:"code"`          // 分类代码

	// 组织信息
	EntityID    *string `json:"entity_id"`    // 实体ID
	IsRecursive *bool   `json:"is_recursive"` // 是否递归
	UserID      *int    `json:"user_id"`      // 用户ID
	GroupID     *int    `json:"group_id"`     // 组ID

	// 分类关联
	ITILCategoryID     *int `json:"itil_category_id"`     // 父级分类ID
	KnowbaseCategoryID *int `json:"knowbase_category_id"` // 知识库分类ID

	// 模板关联
	TicketTemplateIDIncident *int `json:"ticket_template_id_incident"` // 事件工单模板ID
	TicketTemplateIDDemand   *int `json:"ticket_template_id_demand"`   // 需求工单模板ID
	ChangeTemplateID         *int `json:"change_template_id"`          // 变更模板ID
	ProblemTemplateID        *int `json:"problem_template_id"`         // 问题模板ID

	// 类型标识
	IsVisibleIncident *bool `json:"is_visible_incident"` // 是否事件类型
	IsVisibleRequest  *bool `json:"is_visible_request"`  // 是否请求类型
	IsVisibleProblem  *bool `json:"is_visible_problem"`  // 是否问题类型
	IsVisibleChange   *bool `json:"is_visible_change"`   // 是否变更类型
	IsVisibleHelpdesk *bool `json:"is_visible_helpdesk"` // 是否服务台可见

	// 缓存数据
	AncestorCache *string `json:"ancestor_cache"` // 祖先节点缓存
	SonCache      *string `json:"son_cache"`      // 子节点缓存
}

func (m *ITILCategory) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*ITILCategory]())
}
