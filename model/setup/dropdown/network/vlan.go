package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*Vlan]()
}

// Vlan 虚拟局域网
type Vlan struct {
	pkgmodel.Base

	Name *string `json:"name"` // 名称

	// 组织相关
	EntityID    *string `json:"entity_id"`    // 所属实体ID
	IsRecursive *bool   `json:"is_recursive"` // 是否递归继承

	// VLAN配置
	Tag *int `json:"tag"` // VLAN标识(1-4094)
}

func (*Vlan) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*Vlan]())
}
