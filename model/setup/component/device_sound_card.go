package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*DeviceSoundCard]()
}

// DeviceSoundCard 声卡设备
type DeviceSoundCard struct {
	pkgmodel.Base

	Designation *string `json:"designation"` // 声卡标识
	TypeID      *string `json:"type_id"`     // 声卡类型

	// 组织相关
	EntityID    *string `json:"entity_id"`    // 所属实体ID
	IsRecursive *bool   `json:"is_recursive"` // 是否递归继承

	// 制造商信息
	ManufacturerID *string `json:"manufacturer_id"` // 制造商ID
	ModelID        *string `json:"model_id"`        // 型号ID
}

func (*DeviceSoundCard) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_COMPONENT, pkgmodel.GetTableName[*DeviceSoundCard]())
}
