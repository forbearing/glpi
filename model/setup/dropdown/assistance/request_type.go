package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*RequestType]()
}

// RequestType 请求类型
type RequestType struct {
	pkgmodel.Base

	// 基础信息
	Name *string `json:"name"` // 类型名称

	// 状态标识
	IsActive       *bool `json:"is_active"`        // 是否激活
	IsTicketHeader *bool `json:"is_ticket_header"` // 是否工单头部
	IsITILFollowup *bool `json:"is_itil_followup"` // 是否ITIL跟进

	// 默认设置
	IsHelpdeskDefault     *bool `json:"is_helpdesk_default"`      // 是否服务台默认
	IsFollowupDefault     *bool `json:"is_followup_default"`      // 是否跟进默认
	IsMailDefault         *bool `json:"is_mail_default"`          // 是否邮件默认
	IsMailFollowupDefault *bool `json:"is_mail_followup_default"` // 是否邮件跟进默认
}

func (*RequestType) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*RequestType]())
}
