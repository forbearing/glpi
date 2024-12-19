package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*ITILFollowupTemplate]()
}

// ITILFollowupTemplate ITIL跟进模板
type ITILFollowupTemplate struct {
	pkgmodel.Base

	// 基础信息
	Name    *string `json:"name"`    // 模板名称
	Content *string `json:"content"` // 模板内容

	// 组织信息
	EntityID    *string `json:"entity_id"`    // 实体ID
	IsRecursive *bool   `json:"is_recursive"` // 是否递归
	IsPrivate   *bool   `json:"is_private"`   // 是否私有

	// 类型关联
	RequestTypeID *int `json:"request_type_id"` // 请求类型ID
}

func (*ITILFollowupTemplate) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*ITILFollowupTemplate]())
}
