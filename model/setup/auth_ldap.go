package model

import pkgmodel "github.com/forbearing/golib/model"

func init() {
	pkgmodel.Register[*AuthLDAP]()
}

// AuthLDAP LDAP认证配置
type AuthLDAP struct {
	pkgmodel.Base

	Name      *string `json:"name"`       // 配置名称
	IsDefault *bool   `json:"is_default"` // 是否默认
	IsActive  *bool   `json:"is_active"`  // 是否激活

	// 服务器配置
	Host        *string `json:"host"`         // 服务器地址
	Port        *int    `json:"port"`         // 服务器端口
	Timeout     *int    `json:"timeout"`      // 超时时间(秒)
	BaseDN      *string `json:"base_dn"`      // 基础DN
	RootDN      *string `json:"root_dn"`      // 管理员DN
	UseBind     *bool   `json:"use_bind"`     // 是否使用绑定
	UseTLS      *bool   `json:"use_tls"`      // 是否使用TLS
	TLSCertFile *string `json:"tls_certfile"` // TLS证书文件
	TLSKeyFile  *string `json:"tls_keyfile"`  // TLS密钥文件
	TimeOffset  *int    `json:"time_offset"`  // 时间偏移
	DerefOption *int    `json:"deref_option"` // 解引用选项

	// 分页配置
	PageSize           *int  `json:"page_size"`             // 分页大小
	LDAPMaxLimit       *int  `json:"ldap_max_limit"`        // 最大限制
	CanSupportPageSize *bool `json:"can_support_page_size"` // 是否支持分页

	// 搜索条件
	Condition  *string `json:"condition"`   // 搜索条件
	LoginField *string `json:"login_field"` // 登录字段
	SyncField  *string `json:"sync_field"`  // 同步字段
	UseDN      *bool   `json:"use_dn"`      // 是否使用DN

	// 用户字段映射
	RealnameField    *string `json:"realname_field"`    // 真实姓名字段
	FirstnameField   *string `json:"firstname_field"`   // 名字字段
	TitleField       *string `json:"title_field"`       // 职位字段
	CategoryField    *string `json:"category_field"`    // 类别字段
	LanguageField    *string `json:"language_field"`    // 语言字段
	PictureField     *string `json:"picture_field"`     // 头像字段
	LocationField    *string `json:"location_field"`    // 位置字段
	ResponsibleField *string `json:"responsible_field"` // 负责人字段

	// 联系方式字段
	PhoneField  *string `json:"phone_field"`  // 电话字段
	Phone2Field *string `json:"phone2_field"` // 备用电话字段
	MobileField *string `json:"mobile_field"` // 手机字段
	Email1Field *string `json:"email1_field"` // 邮箱1字段
	Email2Field *string `json:"email2_field"` // 邮箱2字段
	Email3Field *string `json:"email3_field"` // 邮箱3字段
	Email4Field *string `json:"email4_field"` // 邮箱4字段

	// 组织字段
	EntityField     *string `json:"entity_field"`              // 实体字段
	EntityCondition *string `json:"entity_condition"`          // 实体条件
	InventoryDomain *string `json:"inventory_domain"`          // 资产域名
	RegNumberField  *string `json:"registration_number_field"` // 注册号字段

	// 组相关
	GroupField       *string `json:"group_field"`        // 组字段
	GroupCondition   *string `json:"group_condition"`    // 组条件
	GroupSearchType  *int    `json:"group_search_type"`  // 组搜索类型
	GroupMemberField *string `json:"group_member_field"` // 组成员字段
}
