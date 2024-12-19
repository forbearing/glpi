package model

import (
	"time"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*Domain]()
}

// Domain 域名信息
type Domain struct {
	pkgmodel.Base

	// 基础信息
	Name      *string    `json:"name"`            // 域名
	IsActive  *bool      `json:"is_active"`       // 是否激活
	DateStart *time.Time `json:"date_start"`      // 域名创建时间
	DateEnd   *time.Time `json:"date_expiration"` // 过期时间

	// 模板相关
	IsTemplate   *bool   `json:"is_template"`   // 是否为模板
	TemplateName *string `json:"template_name"` // 模板名称
	IsRecursive  *bool   `json:"is_recursive"`  // 是否继承父级权限

	// 关联信息
	EntityID     *string `json:"entity_id"`      // 所属实体ID
	DomainTypeID *string `json:"domain_type_id"` // 域名类型ID
	TechUserID   *string `json:"tech_user_id"`   // 技术用户ID
	TechGroupID  *string `json:"tech_group_id"`  // 技术组ID
}
