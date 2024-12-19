package model

import (
	"encoding/json"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*Supplier]()
}

// Supplier 供应商信息
type Supplier struct {
	pkgmodel.Base

	// 基础信息
	Name               *string `json:"name"`                // 供应商名称
	RegistrationNumber *string `json:"registration_number"` // 注册号
	IsActive           *bool   `json:"is_active"`           // 是否激活
	IsRecursive        *bool   `json:"is_recursive"`        // 是否继承父级权限

	// 联系信息
	Address     *string `json:"address"`      // 地址
	Postcode    *string `json:"postcode"`     // 邮编
	Town        *string `json:"town"`         // 城市
	State       *string `json:"state"`        // 州/省
	Country     *string `json:"country"`      // 国家
	Website     *string `json:"website"`      // 网站
	PhoneNumber *string `json:"phone_number"` // 电话
	Fax         *string `json:"fax"`          // 传真
	Email       *string `json:"email"`        // 邮箱

	// 关联信息
	EntityID       *string          `json:"entity_id"`        // 所属实体ID
	SupplierTypeID *string          `json:"supplier_type_id"` // 供应商类型ID
	Picture        *json.RawMessage `json:"picture"`          // 图片信息
}
