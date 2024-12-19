package model

import (
	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*Monitor]()
}

type Monitor struct {
	pkgmodel.Base

	// 基础信息
	Name            *string `json:"name"`             // 显示器名称
	UUID            *string `json:"uuid"`             // 设备唯一标识符
	SerialNumber    *string `json:"serial_number"`    // 硬件序列号
	InventoryNumber *string `json:"inventory_number"` // 资产编号

	// 组织架构信息
	EntityID    *string `json:"entity_id"`    // 所属实体ID
	LocationID  *string `json:"location_id"`  // 物理位置ID
	IsRecursive *bool   `json:"is_recursive"` // 是否继承父级权限
	IsGlobal    *bool   `json:"is_global"`    // 是否为全局设备

	// 配置信息
	ManufacturerID *string  `json:"manufacturer_id"`  // 制造商ID
	MonitorTypeID  *string  `json:"monitor_type_id"`  // 显示器类型ID
	MonitorModelID *string  `json:"monitor_model_id"` // 显示器型号ID
	Size           *float64 `json:"size"`             // 显示器尺寸(英寸)

	// 接口配置
	HasMicrophone  *bool `json:"has_microphone"`  // 是否有麦克风
	HasSpeaker     *bool `json:"has_speaker"`     // 是否有扬声器
	HasVGA         *bool `json:"has_vga"`         // 是否有VGA接口
	HasBNC         *bool `json:"has_bnc"`         // 是否有BNC接口
	HasDVI         *bool `json:"has_dvi"`         // 是否有DVI接口
	HasPivot       *bool `json:"has_pivot"`       // 是否支持旋转
	HasHDMI        *bool `json:"has_hdmi"`        // 是否有HDMI接口
	HasDisplayPort *bool `json:"has_displayport"` // 是否有DisplayPort接口

	// 使用者信息
	UserID        *string `json:"user_id"`        // 主要使用者ID
	GroupID       *string `json:"group_id"`       // 使用者所属组ID
	ContactName   *string `json:"contact_name"`   // 联系人姓名
	ContactNumber *string `json:"contact_number"` // 联系人工号/电话

	// 技术支持信息
	UserTechnicalID  *string `json:"user_technical_id"`  // 负责运维人员ID
	GroupTechnicalID *string `json:"group_technical_id"` // 负责运维团队ID

	// 状态信息
	Status       model.AssetStatus `json:"status"`        // 资产状态
	IsTemplate   *bool             `json:"is_template"`   // 是否为模板记录
	TemplateName *string           `json:"template_name"` // 模板名称

	// 资产同步信息
	IsDynamic          *bool   `json:"is_dynamic"`            // 是否自动更新
	AutoUpdateSystemID *string `json:"auto_update_system_id"` // 更新来源ID

	// 财务信息
	TicketTCO *float64 `json:"ticket_tco"` // 总拥有成本(TCO)
}
