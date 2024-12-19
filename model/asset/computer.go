package model

import (
	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*Computer]()
}

type Computer struct {
	pkgmodel.Base

	// 基础信息
	Name            *string           `json:"name"`             // 计算机名称/主机名
	UUID            *string           `json:"uuid"`             // 设备唯一标识符(BIOS UUID/主板UUID)
	SerialNumber    *string           `json:"serial_number"`    // 硬件序列号(制造商序列号/机器码)
	InventoryNumber *string           `json:"inventory_number"` // 资产编号(企业内部编号)
	Status          model.AssetStatus `json:"status"`           // 资产状态

	// 组织架构信息
	EntityID    *string `json:"entity_id"`    // 所属实体ID(部门/组织架构划分)
	LocationID  *string `json:"location_id"`  // 物理位置ID(办公室/机房等)
	NetworkID   *string `json:"network_id"`   // 网络位置ID(IP地址段/网段划分)
	IsRecursive *bool   `json:"is_recursive"` // 是否继承父级权限

	// 配置信息
	ManufacturerID  *string `json:"manufacturer_id"`   // 制造商ID(联想/戴尔等)
	ComputerModelID *string `json:"computer_model_id"` // 型号ID(具体的产品型号)
	ComputerTypeID  *string `json:"computer_type_id"`  // 类型ID(台式机/笔记本/服务器等)

	// 使用者信息
	UserID        *string `json:"user_id"`        // 主要使用者ID
	GroupID       *string `json:"group_id"`       // 使用者所属组ID
	ContactName   *string `json:"contact_name"`   // 联系人姓名
	ContactNumber *string `json:"contact_number"` // 联系人工号/电话

	// 技术支持信息
	UserTechnicalID  *string `json:"user_technical_id"`  // 负责运维人员ID
	GroupTechnicalID *string `json:"group_technical_id"` // 负责运维团队ID

	// 模版信息
	IsTemplate   *bool   `json:"is_template"`   // 是否为模板记录(用于批量创建)
	TemplateName *string `json:"template_name"` // 模板名称(当IsTemplate为true时使用)

	// 资产同步信息
	IsDynamic          *bool   `json:"is_dynamic"`            // 是否自动更新(true自动/false手动)
	AutoUpdateSystemID string  `json:"auto_update_system_id"` // 更新来源ID(Agent/OCS/手动等)
	LastBoot           *string `json:"last_boot"`             // 最后开机时间(系统启动时间)

	// 财务信息
	TicketTCO *float64 `json:"ticket_tco"` // 总拥有成本(Total Cost of Ownership)
}
