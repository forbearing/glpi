package model

import (
	"encoding/json"
	"time"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*SoftwareLicense]()
}

// SoftwareLicense 软件许可证
type SoftwareLicense struct {
	pkgmodel.Base

	// 基础信息
	Name            *string            `json:"name"`             // 许可证名称
	SerialNumber    *string            `json:"serial_number"`    // 序列号
	InventorySerial *string            `json:"inventory_serial"` // 资产序列号
	Number          *int               `json:"number"`           // 许可证数量
	State           *model.AssetStatus `json:"state"`            // 许可证状态
	IsValid         *bool              `json:"is_valid"`         // 是否有效
	IsRecursive     *bool              `json:"is_recursive"`     // 是否继承父级权限
	IsTemplate      *bool              `json:"is_template"`      // 是否为模板
	TemplateName    *string            `json:"template_name"`    // 模板名称

	// 分层信息
	CompleteName  *string          `json:"complete_name"`  // 完整名称
	Level         *int             `json:"level"`          // 层级
	AncestorCache *json.RawMessage `json:"ancestor_cache"` // 祖先缓存
	SonCache      *json.RawMessage `json:"son_cache"`      // 子级缓存

	// 关联信息
	SoftwareID     *string `json:"software_id"`     // 关联软件ID
	LicenseTypeID  *string `json:"license_type_id"` // 许可证类型ID
	BuyVersionID   *string `json:"buy_version_id"`  // 购买版本ID
	UseVersionID   *string `json:"use_version_id"`  // 使用版本ID
	LocationID     *string `json:"location_id"`     // 位置ID
	ManufacturerID *string `json:"manufacturer_id"` // 制造商ID
	EntityID       *string `json:"entity_id"`       // 所属实体ID

	// 使用限制
	ExpireDate        *time.Time `json:"expire_date"`         // 过期时间
	AllowOverquota    *bool      `json:"allow_overquota"`     // 是否允许超额
	IsHelpdeskVisible *bool      `json:"is_helpdesk_visible"` // 是否在服务台可见

	// 负责人信息
	UserID        *string `json:"user_id"`        // 用户ID
	TechUserID    *string `json:"tech_user_id"`   // 技术负责人ID
	GroupID       *string `json:"group_id"`       // 组ID
	TechGroupID   *string `json:"tech_group_id"`  // 技术组ID
	ContactName   *string `json:"contact_name"`   // 联系人
	ContactNumber *string `json:"contact_number"` // 联系电话

	// 其他
	Picture *json.RawMessage `json:"picture"` // 图片信息
}
