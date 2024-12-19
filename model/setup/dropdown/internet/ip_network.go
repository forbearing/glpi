package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*IPNetwork]()
}

// IPNetwork IP网络配置
type IPNetwork struct {
	pkgmodel.Base

	Name         *string `json:"name"`          // 名称
	CompleteName *string `json:"complete_name"` // 完整名称
	Level        *int    `json:"level"`         // 层级

	// 组织相关
	EntityID      *string `json:"entity_id"`      // 所属实体ID
	IsRecursive   *bool   `json:"is_recursive"`   // 是否递归继承
	ParentD       *int    `json:"parent_id"`      // 父网络ID
	AncestorCache *string `json:"ancestor_cache"` // 祖先缓存
	SonCache      *string `json:"son_cache"`      // 子网缓存

	// IP网络配置
	Version       *int  `json:"version"`        // IP版本(4/6)
	IsAddressable *bool `json:"is_addressable"` // 是否可寻址

	// IPv4地址信息
	Address *string `json:"address"` // IP地址(字符串格式)
	Netmask *string `json:"netmask"` // 子网掩码(字符串格式)
	Gateway *string `json:"gateway"` // 网关地址(字符串格式)

	// IP地址数值表示(用于优化查询)
	Address0 *uint32 `json:"address_0"` // IP地址第0段
	Address1 *uint32 `json:"address_1"` // IP地址第1段
	Address2 *uint32 `json:"address_2"` // IP地址第2段
	Address3 *uint32 `json:"address_3"` // IP地址第3段

	Netmask0 *uint32 `json:"netmask_0"` // 掩码第0段
	Netmask1 *uint32 `json:"netmask_1"` // 掩码第1段
	Netmask2 *uint32 `json:"netmask_2"` // 掩码第2段
	Netmask3 *uint32 `json:"netmask_3"` // 掩码第3段

	Gateway0 *uint32 `json:"gateway_0"` // 网关第0段
	Gateway1 *uint32 `json:"gateway_1"` // 网关第1段
	Gateway2 *uint32 `json:"gateway_2"` // 网关第2段
	Gateway3 *uint32 `json:"gateway_3"` // 网关第3段
}

func (*IPNetwork) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*IPNetwork]())
}
