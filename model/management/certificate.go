package model

import (
	"time"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*Certificate]()
}

// Certificate 证书信息
type Certificate struct {
	pkgmodel.Base

	// 基础信息
	Name   *string            `json:"name"`   // 证书名称
	Status *model.AssetStatus `json:"status"` // 状态ID

	IsRecursive     *bool   `json:"is_recursive"`     // 是否继承父级权限
	IsTemplate      *bool   `json:"is_template"`      // 是否为模板
	IsAutoSign      *bool   `json:"is_auto_sign"`     // 是否自动签名
	TemplateName    *string `json:"template_name"`    // 模板名称
	SerialNumber    *string `json:"serial_number"`    // 序列号
	InventoryNumber *string `json:"inventory_number"` // 其他序列号

	// DNS信息
	DNSName   *string `json:"dns_name"`   // DNS名称
	DNSSuffix *string `json:"dns_suffix"` // DNS后缀

	// 证书内容
	Command            *string    `json:"command"`             // 命令
	CertificateRequest *string    `json:"certificate_request"` // 证书请求
	CertificateItem    *string    `json:"certificate_item"`    // 证书内容
	DateExpiration     *time.Time `json:"date_expiration"`     // 过期时间

	// 联系信息
	ContactName   *string `json:"contact_name"`   // 联系人
	ContactNumber *string `json:"contact_number"` // 联系号码

	// 关联信息
	EntityID          *string `json:"entity_id"`           // 所属实体ID
	CertificateTypeID *string `json:"certificate_type_id"` // 证书类型ID
	UserID            *string `json:"user_id"`             // 用户ID
	GroupID           *string `json:"group_id"`            // 组ID
	TechUserID        *string `json:"tech_user_id"`        // 技术用户ID
	TechGroupID       *string `json:"tech_group_id"`       // 技术组ID
	LocationID        *string `json:"location_id"`         // 位置ID
	ManufacturerID    *string `json:"manufacturer_id"`     // 制造商ID
}
