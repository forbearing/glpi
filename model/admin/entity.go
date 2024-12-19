package model

import pkgmodel "github.com/forbearing/golib/model"

func init() {
	pkgmodel.Register[*Entity]()
}

// Entity GLPI实体信息
type Entity struct {
	pkgmodel.Base

	// 基础信息
	Name          *string `json:"name"`           // 实体名称
	CompleteName  *string `json:"complete_name"`  // 完整名称
	Level         *int    `json:"level"`          // 层级
	EntityID      *string `json:"entity_id"`      // 父实体ID
	AncestorCache *string `json:"ancestor_cache"` // 祖先缓存
	SonCache      *string `json:"son_cache"`      // 子级缓存

	// 注册和标识
	RegistrationNumber *string `json:"registration_number"` // 注册号
	Tag                *string `json:"tag"`                 // 标签

	// 地理位置信息
	Address   *string  `json:"address"`   // 地址
	PostCode  *string  `json:"postcode"`  // 邮政编码
	Town      *string  `json:"town"`      // 城市
	State     *string  `json:"state"`     // 州/省
	Country   *string  `json:"country"`   // 国家
	Latitude  *float64 `json:"latitude"`  // 纬度
	Longitude *float64 `json:"longitude"` // 经度
	Altitude  *float64 `json:"altitude"`  // 海拔

	// 联系信息
	Website     *string `json:"website"`      // 网站
	PhoneNumber *string `json:"phone_number"` // 电话
	Fax         *string `json:"fax"`          // 传真
	Email       *string `json:"email"`        // 邮箱

	// 邮件配置
	AdminEmail       *string `json:"admin_email"`        // 管理员邮箱
	AdminEmailName   *string `json:"admin_email_name"`   // 管理员邮箱显示名称
	FromEmail        *string `json:"from_email"`         // 发件人邮箱
	FromEmailName    *string `json:"from_email_name"`    // 发件人显示名称
	NoReplyEmail     *string `json:"noreply_email"`      // 无需回复邮箱
	NoReplyEmailName *string `json:"noreply_email_name"` // 无需回复邮箱显示名称
	ReplyToEmail     *string `json:"replyto_email"`      // 回复邮箱
	ReplyToEmailName *string `json:"replyto_email_name"` // 回复邮箱显示名称
	MailDomain       *string `json:"mail_domain"`        // 邮件域名
	MailingSignature *string `json:"mailing_signature"`  // 邮件签名

	// LDAP配置
	LDAPDN           *string `json:"ldap_dn"`            // LDAP DN
	AuthLDAPID       *int    `json:"auth_ldap_id"`       // LDAP认证ID
	EntityLDAPFilter *string `json:"entity_ldap_filter"` // 实体LDAP过滤器

	// 告警配置
	CartridgeAlertRepeat            *int  `json:"cartridge_alert_repeat"`               // 墨盒告警重复
	ConsumableAlertRepeat           *int  `json:"consumable_alert_repeat"`              // 耗材告警重复
	UseLicenseAlert                 *bool `json:"use_license_alert"`                    // 启用许可证告警
	SendLicenseAlertBeforeDelay     *int  `json:"send_license_alert_before_delay"`      // 许可证到期提前告警天数
	UseCertificateAlert             *bool `json:"use_certificate_alert"`                // 启用证书告警
	SendCertificateAlertBeforeDelay *int  `json:"send_certificate_alert_before_delay"`  // 证书到期提前告警天数
	CertificateAlertRepeatInterval  *int  `json:"certificate_alert_repeat_interval"`    // 证书告警重复间隔
	UseContractAlert                *bool `json:"use_contract_alert"`                   // 启用合同告警
	SendContractAlertBeforeDelay    *int  `json:"send_contract_alert_before_delay"`     // 合同到期提前告警天数
	UseInfocomAlert                 *bool `json:"use_infocom_alert"`                    // 启用财务信息告警
	SendInfocomAlertBeforeDelay     *int  `json:"send_infocom_alert_before_delay"`      // 财务信息到期提前告警天数
	UseReservationAlert             *bool `json:"use_reservation_alert"`                // 启用预约告警
	UseDomainAlert                  *bool `json:"use_domain_alert"`                     // 启用域名告警
	SendDomainAlertCloseExpiryDelay *int  `json:"send_domain_alert_close_expiry_delay"` // 域名即将过期告警天数
	SendDomainAlertExpiredDelay     *int  `json:"send_domain_alert_expired_delay"`      // 域名过期告警天数

	// 自动化配置
	AutoCloseDelay  *int `json:"auto_close_delay"`  // 自动关闭延迟
	AutoPurgeDelay  *int `json:"auto_purge_delay"`  // 自动清理延迟
	NotClosedDelay  *int `json:"not_closed_delay"`  // 未关闭延迟
	AutoAssignMode  *int `json:"auto_assign_mode"`  // 自动分配模式
	DelaySendEmails *int `json:"delay_send_emails"` // 发送邮件延迟

	// 日历和模板策略
	CalendarStrategy        *int `json:"calendar_strategy"`         // 日历策略
	CalendarID              *int `json:"calendar_id"`               // 日历ID
	TicketTemplateStrategy  *int `json:"ticket_template_strategy"`  // 工单模板策略
	TicketTemplateID        *int `json:"ticket_template_id"`        // 工单模板ID
	ChangeTemplateStrategy  *int `json:"change_template_strategy"`  // 变更模板策略
	ChangeTemplateID        *int `json:"change_template_id"`        // 变更模板ID
	ProblemTemplateStrategy *int `json:"problem_template_strategy"` // 问题模板策略
	ProblemTemplateID       *int `json:"problem_template_id"`       // 问题模板ID

	// 工单和调查配置
	TicketType      *int    `json:"tickettype"`       // 工单类型
	InquestConfig   *int    `json:"inquest_config"`   // 调查配置
	InquestRate     *int    `json:"inquest_rate"`     // 调查率
	InquestDelay    *int    `json:"inquest_delay"`    // 调查延迟
	InquestURL      *string `json:"inquest_URL"`      // 调查URL
	InquestDuration *int    `json:"inquest_duration"` // 调查持续时间

	// 自动填充配置
	AutofillWarrantyDate     *string `json:"autofill_warranty_date"`     // 自动填充保修日期
	AutofillUseDate          *string `json:"autofill_use_date"`          // 自动填充使用日期
	AutofillBuyDate          *string `json:"autofill_buy_date"`          // 自动填充购买日期
	AutofillDeliveryDate     *string `json:"autofill_delivery_date"`     // 自动填充交付日期
	AutofillOrderDate        *string `json:"autofill_order_date"`        // 自动填充订单日期
	AutofillDecommissionDate *string `json:"autofill_decommission_date"` // 自动填充停用日期

	// 默认阈值设置
	DefaultContractAlert            *int `json:"default_contract_alert"`             // 默认合同告警
	DefaultInfocomAlert             *int `json:"default_infocom_alert"`              // 默认财务信息告警
	DefaultCartridgeAlarmThreshold  *int `json:"default_cartridge_alarm_threshold"`  // 默认墨盒告警阈值
	DefaultConsumableAlarmThreshold *int `json:"default_consumable_alarm_threshold"` // 默认耗材告警阈值

	// 其他设置
	IsNotifyEnableDefault  *bool   `json:"is_notify_enable_default"` // 默认启用通知
	SupplierAsPrivate      *bool   `json:"supplier_as_private"`      // 供应商作为私有
	AnonymizeSupportAgents *bool   `json:"anonymize_support_agents"` // 匿名支持代理
	DisplayUserInitial     *bool   `json:"display_user_initial"`     // 显示用户首字母
	EnableCustomCSS        *bool   `json:"enable_custom_css"`        // 启用自定义CSS
	CustomCSSCode          *string `json:"custom_css_code"`          // 自定义CSS代码
	AgentBaseURL           *string `json:"agent_base_url"`           // 代理基础URL

	// 策略设置
	EntityStrategySoftware  *int    `json:"entity_strategy_software"`  // 软件实体策略
	EntityIDSoftware        *string `json:"entity_id_software"`        // 软件实体ID
	ContractStrategyDefault *int    `json:"contract_strategy_default"` // 默认合同策略
	ContractIDDefault       *int    `json:"contract_id_default"`       // 默认合同ID
	TransferStrategy        *int    `json:"transfer_strategy"`         // 转移策略
	TransferID              *int    `json:"transfer_id"`               // 转移ID
}
