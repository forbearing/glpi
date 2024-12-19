package model

import (
	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*SLM]()
}

// SLM 服务级别管理
// SLM(Service Level Management，服务级别管理)是ITIL框架中的重要组成部分，在GLPI中用于管理和监控IT服务水平。
//
// SLA (Service Level Agreement)
// 服务级别协议
// 定义服务质量标准
// 规定响应和解决时间
//
// OLA (Operational Level Agreement)
// 运营级别协议
// 内部支持团队之间的协议
// 定义内部服务标准
//
// SLT (Service Level Target)
// 服务级别目标
// 具体的服务指标
// 可度量的服务标准
type SLM struct {
	pkgmodel.Base

	Name *string `json:"name"` // SLM名称

	// 组织相关
	EntityID    *int  `json:"entity_id"`    // 所属实体ID
	IsRecursive *bool `json:"is_recursive"` // 是否递归继承

	// 日历相关
	UseTicketCalendar *bool `json:"use_ticket_calendar"` // 是否使用工单日历
	CalendarID        *int  `json:"calendar_id"`         // 日历ID
}
