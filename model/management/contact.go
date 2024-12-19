package model

import (
	"encoding/json"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*Contact]()
}

// Contact 联系人信息
type Contact struct {
	pkgmodel.Base

	// 基础信息
	SurName            *string `json:"sur_name"`            // 姓
	FirstName          *string `json:"first_name"`          // 名
	RegistrationNumber *string `json:"registration_number"` // 注册号
	IsRecursive        *bool   `json:"is_recursive"`        // 是否继承父级权限

	// 联系方式
	Phone  *string `json:"phone"`  // 电话
	Phone2 *string `json:"phone2"` // 备用电话
	Mobile *string `json:"mobile"` // 手机
	Fax    *string `json:"fax"`    // 传真
	Email  *string `json:"email"`  // 邮箱

	// 地址信息
	Address  *string `json:"address"`  // 地址
	Postcode *string `json:"postcode"` // 邮编
	Town     *string `json:"town"`     // 城市
	State    *string `json:"state"`    // 州/省
	Country  *string `json:"country"`  // 国家

	// 关联信息
	EntityID      *string          `json:"entity_id"`       // 所属实体ID
	ContactTypeID *string          `json:"contact_type_id"` // 联系人类型ID
	UserTitleID   *string          `json:"user_title_id"`   // 用户头衔ID
	Picture       *json.RawMessage `json:"picture"`         // 图片信息
}
