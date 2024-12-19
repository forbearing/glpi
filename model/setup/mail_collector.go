package model

import pkgmodel "github.com/forbearing/golib/model"

func init() {
	pkgmodel.Register[*MailCollector]()
}

// MailCollector 邮件收集器
type MailCollector struct {
	pkgmodel.Base

	Name        *string `json:"name"`          // 收集器名称
	Host        *string `json:"host"`          // 邮件服务器
	Login       *string `json:"login"`         // 登录账号
	IsActive    *bool   `json:"is_active"`     // 是否激活
	FileSizeMax *int64  `json:"file_size_max"` // 最大文件大小
	UseMailDate *bool   `json:"use_mail_date"` // 使用邮件日期

	// 邮件过滤
	Accepted *string `json:"accepted"` // 接受规则
	Refused  *string `json:"refused"`  // 拒绝规则

	// 收集配置
	RequesterField    *int    `json:"requester_field"`     // 请求者字段
	AddCCToObserver   *bool   `json:"add_cc_to_observer"`  // 是否添加抄送为观察者
	CollectOnlyUnread *bool   `json:"collect_only_unread"` // 仅收集未读
	LastCollectDate   *string `json:"last_collect_date"`   // 上次收集时间

	// 统计信息
	Errors *int `json:"errors"` // 错误次数
}
