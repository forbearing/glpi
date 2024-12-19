package model

import (
	"time"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*Budget]()
}

// Budget 预算信息
type Budget struct {
	pkgmodel.Base

	// 基础信息
	Name        *string  `json:"name"`         // 预算名称
	Value       *float64 `json:"value"`        // 预算金额
	IsRecursive *bool    `json:"is_recursive"` // 是否继承父级权限

	// 模板信息
	IsTemplate   *bool   `json:"is_template"`   // 是否为模板
	TemplateName *string `json:"template_name"` // 模板名称

	// 时间信息
	BeginDate *time.Time `json:"begin_date"` // 开始日期
	EndDate   *time.Time `json:"end_date"`   // 结束日期

	// 关联信息
	EntityID     *string `json:"entity_id"`      // 所属实体ID
	LocationID   *string `json:"location_id"`    // 位置ID
	BudgetTypeID *string `json:"budget_type_id"` // 预算类型ID
}
