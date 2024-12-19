package model

import pkgmodel "github.com/forbearing/golib/model"

type Name struct {
	pkgmodel.Base

	Name *string `json:"name"`
}

type NameEntity struct {
	pkgmodel.Base

	Name        *string `json:"name"`
	EntityID    *string `json:"entity_id"`
	IsRecursive *bool   `json:"is_recursive"`
}

type NameProductNumber struct {
	pkgmodel.Base

	Name          *string `json:"name"`           // 型号名称
	ProductNumber *string `json:"product_number"` // 产品编号
}
