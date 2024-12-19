package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*DeviceProcessor]()
}

// DeviceProcessor 处理器设备
type DeviceProcessor struct {
	pkgmodel.Base

	Designation *string `json:"designation"` // 处理器标识

	// 组织相关
	EntityID    *string `json:"entity_id"`    // 所属实体ID
	IsRecursive *bool   `json:"is_recursive"` // 是否递归继承

	// 制造商信息
	ManufacturerID *string `json:"manufacturer_id"` // 制造商ID
	ModelID        *string `json:"model_id"`        // 型号ID

	// 规格参数
	Frequency        *int64 `json:"frequence"`         // 当前频率(MHz)
	FrequencyDefault *int64 `json:"frequency_default"` // 默认频率(MHz)
	CoreCount        *int   `json:"cores_count"`       // 核心数量
	ThreadCount      *int   `json:"thread_count"`      // 线程数量
}

func (*DeviceProcessor) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_COMPONENT, pkgmodel.GetTableName[*DeviceProcessor]())
}
