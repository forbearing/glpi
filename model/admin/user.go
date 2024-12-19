package model

import (
	"time"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*User]()
}

// User GLPI用户信息
type User struct {
	pkgmodel.Base

	// 基本个人信息
	Name               *string `json:"name"`                // 用户名
	RealName           *string `json:"real_name"`           // 真实姓名
	FirstName          *string `json:"first_name"`          // 名字
	Nickname           *string `json:"nickname"`            // 昵称
	RegistrationNumber *string `json:"registration_number"` // 注册号
	Picture            *string `json:"picture"`             // 头像
	Phone              *string `json:"phone"`               // 电话
	Phone2             *string `json:"phone2"`              // 备用电话
	Mobile             *string `json:"mobile"`              // 手机号

	// 账户状态
	IsActive       *bool      `json:"is_active"`        // 是否激活
	IsDeletedLDAP  *bool      `json:"is_deleted_ldap"`  // LDAP是否已删除
	LastLogin      *time.Time `json:"last_login"`       // 最后登录时间
	ValidBeginDate *time.Time `json:"valid_begin_date"` // 账户生效时间
	ValidEndDate   *time.Time `json:"valid_end_date"`   // 账户过期时间

	// 认证相关
	AuthType                *int       `json:"auth_type"`                  // 认证类型
	AuthID                  *int       `json:"auth_id"`                    // 认证ID
	PasswordLastUpdate      *time.Time `json:"password_last_update"`       // 密码最后更新时间
	PasswordForgetToken     *string    `json:"password_forget_token"`      // 忘记密码令牌
	PasswordForgetTokenDate *time.Time `json:"password_forget_token_date"` // 忘记密码令牌日期
	PersonalTokenDate       *time.Time `json:"personal_token_date"`        // 个人令牌日期
	APITokenDate            *time.Time `json:"api_token_date"`             // API令牌日期
	CookieTokenDate         *time.Time `json:"cookie_token_date"`          // Cookie令牌日期
	UserDN                  *string    `json:"user_dn"`                    // LDAP DN
	UserDNHash              *string    `json:"user_dn_hash"`               // LDAP DN哈希

	// 界面偏好设置
	Language        *string `json:"language"`         // 语言偏好
	UseMode         *int    `json:"use_mode"`         // 使用模式
	ListLimit       *int    `json:"list_limit"`       // 列表限制数
	DateFormat      *string `json:"date_format"`      // 日期格式
	NumberFormat    *string `json:"number_format"`    // 数字格式
	NameFormat      *string `json:"name_format"`      // 名称格式
	CSVDelimiter    *string `json:"csv_delimiter"`    // CSV分隔符
	PDFFont         *string `json:"pdf_font"`         // PDF字体
	HighContrastCSS *bool   `json:"highcontrast_css"` // 高对比度CSS
	Palette         *string `json:"palette"`          // 调色板
	PageLayout      *string `json:"page_layout"`      // 页面布局
	FoldMenu        *bool   `json:"fold_menu"`        // 折叠菜单
	FoldSearch      *bool   `json:"fold_search"`      // 折叠搜索

	// 显示设置
	IsIDVisible                       *bool `json:"is_id_visible"`                          // 是否显示ID
	UseFlatDropdownTree               *bool `json:"use_flat_dropdowntree"`                  // 使用平面下拉树
	UseFlatDropdownTreeOnSearchResult *bool `json:"use_flat_dropdowntree_on_search_result"` // 搜索结果使用平面下拉树
	ShowJobAtLogin                    *bool `json:"show_job_at_login"`                      // 登录时显示工作
	ShowCountOnTabs                   *bool `json:"show_count_on_tabs"`                     // 标签上显示计数
	DisplayCountOnHome                *bool `json:"display_count_on_home"`                  // 主页显示计数
	RefreshViews                      *bool `json:"refresh_views"`                          // 刷新视图

	// 任务和提醒相关
	TaskState               *string `json:"task_state"`                 // 任务状态
	TaskPrivate             *bool   `json:"task_private"`               // 任务私有
	FollowupPrivate         *bool   `json:"followup_private"`           // 跟进私有
	TimelineOrder           *string `json:"timeline_order"`             // 时间线顺序
	TimelineActionBtnLayout *string `json:"timeline_action_btn_layout"` // 时间线动作按钮布局
	TimelineDateFormat      *string `json:"timeline_date_format"`       // 时间线日期格式

	// 优先级设置
	Priority1 *int `json:"priority_1"` // 优先级1
	Priority2 *int `json:"priority_2"` // 优先级2
	Priority3 *int `json:"priority_3"` // 优先级3
	Priority4 *int `json:"priority_4"` // 优先级4
	Priority5 *int `json:"priority_5"` // 优先级5
	Priority6 *int `json:"priority_6"` // 优先级6

	// 截止日期颜色设置
	DueDateOKColor       *string `json:"duedate_ok_color"`       // 正常截止日期颜色
	DueDateWarningColor  *string `json:"duedate_warning_color"`  // 警告截止日期颜色
	DueDateCriticalColor *string `json:"duedate_critical_color"` // 危险截止日期颜色
	DueDateWarningLess   *int    `json:"duedate_warning_less"`   // 警告截止日期阈值
	DueDateCriticalLess  *int    `json:"duedate_critical_less"`  // 危险截止日期阈值
	DueDateWarningUnit   *string `json:"duedate_warning_unit"`   // 警告截止日期单位
	DueDateCriticalUnit  *string `json:"duedate_critical_unit"`  // 危险截止日期单位

	// 布局和配置JSON
	ITILLayout           *string `json:"itil_layout"`            // ITIL布局配置
	RichtextLayout       *string `json:"richtext_layout"`        // 富文本布局配置
	Plannings            *string `json:"plannings"`              // 计划配置
	DisplayOptions       *string `json:"display_options"`        // 显示选项
	PrivateBookmarkOrder *string `json:"private_bookmark_order"` // 私有书签顺序
	SavedSearchPinned    *string `json:"saved_search_pinned"`    // 固定的已保存搜索

	// 默认设置
	DefaultRequestTypeID       *int    `json:"default_request_type_id"`       // 默认请求类型ID
	DefaultDashboardCentral    *string `json:"default_dashboard_central"`     // 默认中央仪表板
	DefaultDashboardAsset      *string `json:"default_dashboard_asset"`       // 默认资产仪表板
	DefaultDashboardHelpdesk   *string `json:"default_dashboard_helpdesk"`    // 默认服务台仪表板
	DefaultDashboardMiniTicket *string `json:"default_dashboard_mini_ticket"` // 默认迷你工单仪表板
	DefaultCentralTab          *string `json:"default_central_tab"`           // 默认中央标签

	// 关联ID
	LocationID       *string `json:"location_id"`        // 位置ID
	ProfileID        *string `json:"profile_id"`         // 配置文件ID
	EntityID         *string `json:"entity_id"`          // 实体ID
	UserTitleID      *string `json:"user_title_id"`      // 用户头衔ID
	UserCategoryID   *string `json:"user_category_id"`   // 用户分类ID
	GroupID          *string `json:"group_id"`           // 组ID
	UserSupervisorID *string `json:"user_supervisor_id"` // 主管用户ID

	// 其他设置
	SetDefaultTech               *bool   `json:"set_default_tech"`                // 设置默认技术人员
	SetDefaultRequester          *bool   `json:"set_default_requester"`           // 设置默认请求者
	NotificationToMyself         *bool   `json:"notification_to_myself"`          // 向自己发送通知
	BackCreated                  *bool   `json:"back_created"`                    // 返回已创建
	KeepDeviceWhenPurgingItem    *bool   `json:"keep_device_when_purging_item"`   // 清除项目时保留设备
	LockAutolockMode             *string `json:"lock_auto_lock_mode"`             // 锁定自动锁定模式
	LockDirectunlockNotification *bool   `json:"lock_direct_unlock_notification"` // 锁定直接解锁通知
	Timezone                     *string `json:"timezone"`                        // 时区
	SyncField                    *string `json:"sync_field"`                      // 同步字段
}
