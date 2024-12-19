package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*Location]()
}

// Location 位置信息
type Location struct {
	pkgmodel.Base

	// 基础信息
	Name         *string `json:"name"`          // 位置名称
	CompleteName *string `json:"complete_name"` // 完整位置名称
	Level        *int    `json:"level"`         // 层级

	// 组织信息
	EntityID    *string `json:"entity_id"`    // 实体ID
	IsRecursive *bool   `json:"is_recursive"` // 是否递归
	ParentID    *int    `json:"parent_id"`    // 父级位置ID

	// 层级缓存
	AncestorCache *string `json:"ancestor_cache"` // 祖先节点缓存
	SonCache      *string `json:"son_cache"`      // 子节点缓存

	// 地址信息
	Address  *string `json:"address"`  // 详细地址
	Postcode *string `json:"postcode"` // 邮政编码
	Town     *string `json:"town"`     // 城镇
	State    *string `json:"state"`    // 州/省
	Country  *string `json:"country"`  // 国家
	Building *string `json:"building"` // 建筑
	Room     *string `json:"room"`     // 房间

	// 地理坐标
	Latitude  *string `json:"latitude"`  // 纬度
	Longitude *string `json:"longitude"` // 经度
	Altitude  *string `json:"altitude"`  // 海拔
}

func (*Location) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*Location]())
}
