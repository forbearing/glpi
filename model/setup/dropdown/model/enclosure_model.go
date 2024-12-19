package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*EnclosureModel]()
}

// EnclosureModel 机箱型号
type EnclosureModel struct {
	pkgmodel.Base

	// 基础信息
	Name          *string `json:"name"`           // 名称
	ProductNumber *string `json:"product_number"` // 产品编号

	// 物理特性
	Weight       *int  `json:"weight"`        // 重量
	RequiredUnit *int  `json:"required_unit"` // 所需机架单元数
	Depth        *int  `json:"depth"`         // 深度
	IsHalfRack   *bool `json:"is_half_rack"`  // 是否为半机架

	// 电源特性
	PowerConnection  *int `json:"power_connection"`  // 电源连接数
	PowerConsumption *int `json:"power_consumption"` // 功耗

	// 图片资源
	PictureFront *string `json:"picture_front"` // 前视图
	PictureRear  *string `json:"picture_rear"`  // 后视图
	Picture      *string `json:"picture"`       // 其他图片
}

func (*EnclosureModel) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*EnclosureModel]())
}
