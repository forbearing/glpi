package model

import pkgmodel "github.com/forbearing/golib/model"

func init() {
	pkgmodel.Register[*Group]()
}

// Group 用户组
type Group struct {
	pkgmodel.Base

	// 基础信息
	Name         *string `json:"name"`          // 组名称
	CompleteName *string `json:"complete_name"` // 完整名称
	Level        *int    `json:"level"`         // 层级
	IsRecursive  *bool   `json:"is_recursive"`  // 是否递归

	// 权限标记
	IsRequester *bool `json:"is_requester"`  // 是否可以提交请求
	IsWatcher   *bool `json:"is_watcher"`    // 是否可以查看
	IsAssign    *bool `json:"is_assign"`     // 是否可以分配
	IsTask      *bool `json:"is_task"`       // 是否可以处理任务
	IsNotify    *bool `json:"is_notify"`     // 是否接收通知
	IsItemGroup *bool `json:"is_item_group"` // 是否为项目组
	IsUserGroup *bool `json:"is_user_group"` // 是否为用户组
	IsManager   *bool `json:"is_manager"`    // 是否为管理组

	// LDAP相关
	LDAPField   *string `json:"ldap_field"`    // LDAP字段
	LDAPValue   *string `json:"ldap_value"`    // LDAP值
	LDAPGroupDN *string `json:"ldap_group_dn"` // LDAP组DN

	// 层级关系
	GroupID       *int    `json:"group_id"`       // 父组ID
	AncestorCache *string `json:"ancestor_cache"` // 祖先缓存
	SonCache      *string `json:"son_cache"`      // 子级缓存

	// 关联信息
	EntityID *string `json:"entity_id"` // 实体ID
}
