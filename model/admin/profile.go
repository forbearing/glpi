package model

import pkgmodel "github.com/forbearing/golib/model"

func init() {
	pkgmodel.Register[*Profile]()
}

// Profile 配置文件信息
type Profile struct {
	pkgmodel.Base

	// 基础信息
	Name      *string `json:"name"`       // 配置文件名称
	Interface *string `json:"interface"`  // 界面类型(如: helpdesk)
	IsDefault *bool   `json:"is_default"` // 是否为默认配置

	// 服务台配置
	HelpdeskHardware *bool   `json:"helpdesk_hardware"`  // 是否启用服务台硬件支持
	HelpdeskItemType *string `json:"helpdesk_item_type"` // 服务台可处理的设备类型

	// 工单相关
	TicketStatus        *string `json:"ticket_status"`          // 工单状态配置
	CreateTicketOnLogin *bool   `json:"create_ticket_on_login"` // 登录时是否创建工单
	TicketTemplateID    *int    `json:"ticket_template_id"`     // 工单模板ID

	// 变更相关
	ChangeStatus     *string `json:"change_status"`      // 变更状态配置
	ChangeTemplateID *int    `json:"change_template_id"` // 变更模板ID

	// 问题相关
	ProblemStatus     *string `json:"problem_status"`      // 问题状态配置
	ProblemTemplateID *int    `json:"problem_template_id"` // 问题模板ID

	// DNS配置
	ManagedDomainRecordTypes *string `json:"managed_domain_record_types"` // 可管理的DNS记录类型
}
