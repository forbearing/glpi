package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*PCIVendor]()
}

// PCIVendor USB供应商信息
type PCIVendor struct {
	pkgmodel.Base

	Name *string `json:"name"` // 名称

	// 组织相关
	EntityID    *int  `json:"entity_id"`    // 所属实体ID
	IsRecursive *bool `json:"is_recursive"` // 是否递归继承

	// USB标识
	VendorID *string `json:"vendor_id"` // 供应商ID
	DeviceID *string `json:"device_id"` // 设备ID
}

func (*PCIVendor) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*PCIVendor]())
}
