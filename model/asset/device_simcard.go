package model

import (
	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*DeviceSimcard]()
	pkgmodel.Register[*DeviceSimcardComponent]()
}

// DeviceSimcard 设备SIM卡关联信息
type DeviceSimcard struct {
	pkgmodel.Base

	// 关联信息
	DeviceID    *string `json:"device_id"`    // 关联设备ID
	DeviceType  *string `json:"device_type"`  // 关联设备类型
	ComponentID *string `json:"component_id"` // SIM卡型号ID

	// 基础信息
	SerialNumber    *string           `json:"serial_number"`    // 序列号
	InventoryNumber *string           `json:"inventory_number"` // 资产编号
	Status          model.AssetStatus `json:"status"`           // 状态
	IsDynamic       *bool             `json:"is_dynamic"`       // 是否动态

	// 组织架构信息
	EntityID    *string `json:"entity_id"`    // 所属实体ID
	LocationID  *string `json:"location_id"`  // 物理位置ID
	IsRecursive *bool   `json:"is_recursive"` // 是否继承父级权限
	UserID      *string `json:"user_id"`      // 使用者ID
	GroupID     *string `json:"group_id"`     // 使用组ID

	// SIM卡特有信息
	PIN    *string `json:"pin"`     // PIN码
	PIN2   *string `json:"pin2"`    // PIN2码
	PUK    *string `json:"puk"`     // PUK码
	PUK2   *string `json:"puk2"`    // PUK2码
	MSIN   *string `json:"msin"`    // 移动用户识别码
	LineID *string `json:"line_id"` // 通信线路ID
}

// DeviceSimcardComponent SIM卡设备组件
type DeviceSimcardComponent struct {
	pkgmodel.Base

	// 基础信息
	Name      *string `json:"name"`       // 组件名称
	Voltage   *int64  `json:"voltage"`    // 工作电压(mV)
	AllowVoIP *bool   `json:"allow_voip"` // 是否支持VoIP

	// 组织架构信息
	EntityID    *string `json:"entity_id"`    // 所属实体ID
	IsRecursive *bool   `json:"is_recursive"` // 是否继承父级权限

	// 配置信息
	ManufacturerID *string `json:"manufacturer_id"` // 制造商ID
	SimcardTypeID  *string `json:"simcard_type_id"` // SIM卡类型ID
}
