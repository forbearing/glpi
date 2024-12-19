package model

import (
	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*Software]()
}

type Software struct {
	pkgmodel.Base

	// 基础信息
	Name     *string           `json:"name"`      // 软件名称
	Status   model.AssetStatus `json:"status"`    // 软件状态(is_valid)
	Picture  *string           `json:"picture"`   // 软件图片
	ParentID *string           `json:"parent_id"` // 父级软件ID(softwares_id)

	// 组织架构信息
	EntityID    *string `json:"entity_id"`    // 所属实体ID
	LocationID  *string `json:"location_id"`  // 物理位置ID
	IsRecursive *bool   `json:"is_recursive"` // 是否继承父级权限

	// 分类信息
	ManufacturerID *string `json:"manufacturer_id"` // 制造商ID
	CategoryID     *string `json:"category_id"`     // 软件分类ID

	// 使用者信息
	UserID  *string `json:"user_id"`  // 主要使用者ID
	GroupID *string `json:"group_id"` // 使用者所属组ID

	// 技术支持信息
	UserTechnicalID  *string `json:"user_technical_id"`  // 负责运维人员ID
	GroupTechnicalID *string `json:"group_technical_id"` // 负责运维团队ID

	// 状态信息
	IsUpdate          *bool   `json:"is_update"`           // 是否为更新版本
	IsTemplate        *bool   `json:"is_template"`         // 是否为模板记录
	TemplateName      *string `json:"template_name"`       // 模板名称
	IsHelpdeskVisible *bool   `json:"is_helpdesk_visible"` // 是否在服务台可见

	// 财务信息
	TicketTCO *float64 `json:"ticket_tco"` // 总拥有成本(TCO)
}
