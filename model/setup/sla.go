package model

import pkgmodel "github.com/forbearing/golib/model"

func init() {
	pkgmodel.Register[*SLA]()
}

// SLA 服务级别协议
type SLA struct {
	pkgmodel.Base

	Name *string `json:"name"` // SLA名称

	// 组织相关
	EntityID    *string `json:"entity_id"`    // 所属实体ID
	IsRecursive *bool   `json:"is_recursive"` // 是否递归继承

	// SLM相关
	SLMID *string `json:"slm_id"` // 所属SLM ID
	Type  *int    `json:"type"`   // SLA类型

	// 时间相关
	NumberTime        *int    `json:"number_time"`         // 时间数值
	DefinitionTime    *string `json:"definition_time"`     // 时间单位(hour/minute等)
	UseTicketCalendar *bool   `json:"use_ticket_calendar"` // 是否使用工单日历
	CalendarID        *string `json:"calendar_id"`         // 日历ID
	EndOfWorkingDay   *bool   `json:"end_of_working_day"`  // 是否工作日结束
}
