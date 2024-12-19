package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*DevicePCI]()
}

// DevicePCI PCI设备
type DevicePCI struct {
	pkgmodel.Base

	Designation *string `json:"designation"` // PCI设备标识

	// 组织相关
	EntityID    *string `json:"entity_id"`    // 所属实体ID
	IsRecursive *bool   `json:"is_recursive"` // 是否递归继承

	// 制造商信息
	ManufacturerID     *string `json:"manufacturer_id"`       // 制造商ID
	ModelID            *string `json:"model_id"`              // PCI型号ID
	NetworkCardModelID *string `json:"network_card_model_id"` // 网卡型号ID(如果是网卡)
}

func (*DevicePCI) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_COMPONENT, pkgmodel.GetTableName[*DevicePCI]())
}
