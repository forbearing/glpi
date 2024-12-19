package model

import (
	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*Cable]()
}

// Cable 线缆资产
type Cable struct {
	pkgmodel.Base

	// 基础信息
	Name            *string           `json:"name"`             // 线缆名称
	InventoryNumber *string           `json:"inventory_number"` // 资产编号
	Status          model.AssetStatus `json:"status"`           // 资产状态
	Color           *string           `json:"color"`            // 线缆颜色

	// 端点A信息
	EndpointAType        *string `json:"endpoint_a_type"`         // A端设备类型
	EndpointAItemID      *string `json:"endpoint_a_item_id"`      // A端设备ID
	EndpointASocket      *string `json:"endpoint_a_socket"`       // A端接口
	EndpointASocketModel *string `json:"endpoint_a_socket_model"` // A端接口型号

	// 端点B信息
	EndpointBType        *string `json:"endpoint_b_type"`         // B端设备类型
	EndpointBItemID      *string `json:"endpoint_b_item_id"`      // B端设备ID
	EndpointBSocket      *string `json:"endpoint_b_socket"`       // B端接口
	EndpointBSocketModel *string `json:"endpoint_b_socket_model"` // B端接口型号

	// 组织架构信息
	EntityID    *string `json:"entity_id"`    // 所属实体ID
	IsRecursive *bool   `json:"is_recursive"` // 是否继承父级权限

	// 配置信息
	CableTypeID   *string `json:"cable_type_id"`   // 线缆类型ID
	CableStrandID *string `json:"cable_strand_id"` // 线缆芯数ID

	// 技术支持信息
	UserTechnicalID *string `json:"user_technical_id"` // 负责运维人员ID
}
