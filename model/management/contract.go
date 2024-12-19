package model

import (
	"time"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

// Contract 合同
type Contract struct {
	pkgmodel.Base

	// 基础信息
	Name             *string            `json:"name"`              // 名称
	Number           *string            `json:"number"`            // 编号
	EntityID         *string            `json:"entity_id"`         // 实体ID
	IsRecursive      *bool              `json:"is_recursive"`      // 是否递归
	ContractTypeID   *int               `json:"contract_type_id"`  // 合同类型ID
	LocationID       *int               `json:"location_id"`       // 位置ID
	AccountingNumber *string            `json:"accounting_number"` // 会计编号
	Status           *model.AssetStatus `json:"status"`            // 状态ID

	// 时间相关
	BeginDate   *time.Time `json:"begin_date"`  // 开始日期
	Duration    *int       `json:"duration"`    // 持续时间
	Notice      *int       `json:"notice"`      // 通知期
	Periodicity *int       `json:"periodicity"` // 周期
	Billing     *int       `json:"billing"`     // 计费周期
	Alert       *int       `json:"alert"`       // 警报
	Renewal     *int       `json:"renewal"`     // 续约

	// 工作时间
	WeekBeginHour     *string `json:"week_begin_hour"`     // 工作日开始时间
	WeekEndHour       *string `json:"week_end_hour"`       // 工作日结束时间
	SaturdayBeginHour *string `json:"saturday_begin_hour"` // 周六开始时间
	SaturdayEndHour   *string `json:"saturday_end_hour"`   // 周六结束时间
	UseSaturday       *bool   `json:"use_saturday"`        // 是否使用周六
	SundayBeginHour   *string `json:"sunday_begin_hour"`   // 周日开始时间
	SundayEndHour     *string `json:"sunday_end_hour"`     // 周日结束时间
	UseSunday         *bool   `json:"use_sunday"`          // 是否使用周日

	// 模板相关
	IsTemplate     *bool   `json:"is_template"`      // 是否为模板
	TemplateName   *string `json:"template_name"`    // 模板名称
	MaxLinkAllowed *int    `json:"max_link_allowed"` // 允许的最大链接数
}
