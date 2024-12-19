package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*PhoneModel]()
}

// PhoneModel 电话型号
type PhoneModel struct {
	pkgmodel.Base

	// 基础信息
	Name          *string `json:"name"`           // 型号名称
	ProductNumber *string `json:"product_number"` // 产品编号

	// 图片信息
	PictureFront *string `json:"picture_front"` // 前视图
	PictureRear  *string `json:"picture_rear"`  // 后视图
	Picture      *string `json:"picture"`       // 其他图片
}

func (*PhoneModel) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*PhoneModel]())
}
