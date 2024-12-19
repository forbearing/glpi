package model

import (
	"fmt"
	"time"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*Holiday]()
}

// Holiday 假期/节假日
type Holiday struct {
	pkgmodel.Base

	Name *string `json:"name"` // 名称

	// 组织相关
	EntityID    *string `json:"entity_id"`    // 所属实体ID
	IsRecursive *bool   `json:"is_recursive"` // 是否递归继承

	// 时间设置
	BeginDate   *time.Time `json:"begin_date"`   // 开始日期
	EndDate     *time.Time `json:"end_date"`     // 结束日期
	IsPerpetual *bool      `json:"is_perpetual"` // 是否永久性假期
}

func (*Holiday) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*Holiday]())
}
