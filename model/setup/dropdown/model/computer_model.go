package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*ComputerModel]()
}

// ComputerModel 计算机型号
type ComputerModel struct {
	pkgmodel.Base

	// 基础信息
	Name          *string `json:"name"`           // 型号名称
	ProductNumber *string `json:"product_number"` // 产品编号

	// 物理特征
	Weight       *float64 `json:"weight"`        // 重量
	Depth        *float64 `json:"depth"`         // 深度
	RequiredUnit *int     `json:"required_unit"` // 所需单位
	IsHalfRack   *bool    `json:"is_half_rack"`  // 是否半机架

	// 电源信息
	PowerConnection  *int `json:"power_connection"`  // 电源接口数量
	PowerConsumption *int `json:"power_consumption"` // 功率消耗

	// 图片信息
	PictureFront *string `json:"picture_front"` // 前视图
	PictureRear  *string `json:"picture_rear"`  // 后视图
	Picture      *string `json:"picture"`       // 其他图片
}

func (*ComputerModel) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*ComputerModel]())
}
