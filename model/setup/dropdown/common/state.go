package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*Status]()
}

// Status 状态信息
type Status struct {
	pkgmodel.Base

	// 基础信息
	Name         *string `json:"name"`          // 状态名称
	CompleteName *string `json:"complete_name"` // 完整状态名称
	Level        *int    `json:"level"`         // 层级

	// 组织信息
	EntityID    *string `json:"entity_id"`    // 实体ID
	IsRecursive *bool   `json:"is_recursive"` // 是否递归
	ParentID    *int    `json:"parent_id"`    // 父级状态ID

	// 层级缓存
	AncestorCache *string `json:"ancestor_cache"` // 祖先节点缓存
	SonCache      *string `json:"son_cache"`      // 子节点缓存

	// 可见性配置
	IsVisibleComputer           *bool `json:"is_visible_computer"`           // 计算机可见性
	IsVisibleMonitor            *bool `json:"is_visible_monitor"`            // 显示器可见性
	IsVisibleNetworkequipment   *bool `json:"is_visible_networkequipment"`   // 网络设备可见性
	IsVisiblePeripheral         *bool `json:"is_visible_peripheral"`         // 外围设备可见性
	IsVisiblePhone              *bool `json:"is_visible_phone"`              // 电话可见性
	IsVisiblePrinter            *bool `json:"is_visible_printer"`            // 打印机可见性
	IsVisibleSoftwareversion    *bool `json:"is_visible_softwareversion"`    // 软件版本可见性
	IsVisibleSoftwarelicense    *bool `json:"is_visible_softwarelicense"`    // 软件许可可见性
	IsVisibleLine               *bool `json:"is_visible_line"`               // 线路可见性
	IsVisibleCertificate        *bool `json:"is_visible_certificate"`        // 证书可见性
	IsVisibleRack               *bool `json:"is_visible_rack"`               // 机架可见性
	IsVisiblePassivedcequipment *bool `json:"is_visible_passivedcequipment"` // 被动DC设备可见性
	IsVisibleEnclosure          *bool `json:"is_visible_enclosure"`          // 机柜可见性
	IsVisiblePdu                *bool `json:"is_visible_pdu"`                // PDU可见性
	IsVisibleCluster            *bool `json:"is_visible_cluster"`            // 集群可见性
	IsVisibleContract           *bool `json:"is_visible_contract"`           // 合同可见性
	IsVisibleAppliance          *bool `json:"is_visible_appliance"`          // 设备可见性
	IsVisibleDatabaseinstance   *bool `json:"is_visible_databaseinstance"`   // 数据库实例可见性
	IsVisibleCable              *bool `json:"is_visible_cable"`              // 线缆可见性
	IsVisibleUnmanaged          *bool `json:"is_visible_unmanaged"`          // 非托管设备可见性
}

func (*Status) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*Status]())
}
