package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*DeviceNetworkCard]()
}

// DeviceNetworkCard 网卡设备
type DeviceNetworkCard struct {
	pkgmodel.Base

	Designation *string `json:"designation"` // 网卡标识

	// 组织相关
	EntityID    *string `json:"entity_id"`    // 所属实体ID
	IsRecursive *bool   `json:"is_recursive"` // 是否递归继承

	// 制造商信息
	ManufacturerID *string `json:"manufacturer_id"` // 制造商ID
	ModelID        *string `json:"model_id"`        // 型号ID

	// 网络参数
	Bandwidth  *int64  `json:"bandwidth"`   // 带宽
	MACDefault *string `json:"mac_default"` // 默认MAC地址
}

func (*DeviceNetworkCard) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_COMPONENT, pkgmodel.GetTableName[*DeviceNetworkCard]())
}
