package model

import (
	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

type Printer struct {
	pkgmodel.Base

	// 基础信息
	Name            *string           `json:"name"`             // 打印机名称
	UUID            *string           `json:"uuid"`             // 设备唯一标识符
	Status          model.AssetStatus `json:"status"`           // 资产状态
	SerialNumber    *string           `json:"serial_number"`    // 硬件序列号
	InventoryNumber *string           `json:"inventory_number"` // 资产编号
	SysDescription  *string           `json:"sys_description"`  // 系统描述

	// 组织架构信息
	EntityID    *string `json:"entity_id"`    // 所属实体ID
	LocationID  *string `json:"location_id"`  // 物理位置ID
	NetworkID   *string `json:"network_id"`   // 网络位置ID
	IsRecursive *bool   `json:"is_recursive"` // 是否继承父级权限
	IsGlobal    *bool   `json:"is_global"`    // 是否为全局设备

	// 硬件信息
	MemorySize       *float64 `json:"memory_size"`        // 内存大小
	HasSerial        *bool    `json:"has_serial"`         // 是否有串口
	HasParallel      *bool    `json:"has_parallel"`       // 是否有并口
	HasUSB           *bool    `json:"has_usb"`            // 是否有USB接口
	HasWifi          *bool    `json:"has_wifi"`           // 是否有无线网络
	HasEthernet      *bool    `json:"has_ethernet"`       // 是否有有线网络
	SNMPCredentialID *string  `json:"snmp_credential_id"` // SNMP凭证ID

	// 配置信息
	ManufacturerID *string `json:"manufacturer_id"`  // 制造商ID
	PrinterTypeID  *string `json:"printer_type_id"`  // 打印机类型ID
	PrinterModelID *string `json:"printer_model_id"` // 打印机型号ID

	// 使用者信息
	UserID        *string `json:"user_id"`        // 主要使用者ID
	GroupID       *string `json:"group_id"`       // 使用者所属组ID
	ContactName   *string `json:"contact_name"`   // 联系人姓名
	ContactNumber *string `json:"contact_number"` // 联系人工号/电话

	// 技术支持信息
	UserTechnicalID  *string `json:"user_technical_id"`  // 负责运维人员ID
	GroupTechnicalID *string `json:"group_technical_id"` // 负责运维团队ID

	// 模版信息
	IsTemplate          *bool   `json:"is_template"`           // 是否为模板记录
	TemplateName        *string `json:"template_name"`         // 模板名称
	LastInventoryUpdate *string `json:"last_inventory_update"` // 最后清单更新时间

	// 计数器信息
	InitPagesCounter *int64 `json:"init_pages_counter"` // 初始页数计数
	LastPagesCounter *int64 `json:"last_pages_counter"` // 最新页数计数

	// 资产同步信息
	IsDynamic          *bool   `json:"is_dynamic"`            // 是否自动更新
	AutoUpdateSystemID *string `json:"auto_update_system_id"` // 更新来源ID

	// 财务信息
	TicketTCO *float64 `json:"ticket_tco"` // 总拥有成本(TCO)
}
