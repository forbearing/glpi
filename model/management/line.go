package model

import (
	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*Line]()
}

// Line 线路信息
type Line struct {
	pkgmodel.Base

	// 基础信息
	Name        *string `json:"name"`         // 线路名称
	IsRecursive *bool   `json:"is_recursive"` // 是否继承父级权限

	// 呼叫信息
	CallerNumber *string `json:"caller_number"` // 呼叫号码
	CallerName   *string `json:"caller_name"`   // 呼叫者名称

	// 关联信息
	EntityID       *string            `json:"entity_id"`        // 所属实体ID
	UserID         *string            `json:"user_id"`          // 用户ID
	GroupID        *string            `json:"group_id"`         // 组ID
	LineOperatorID *string            `json:"line_operator_id"` // 线路运营商ID
	LocationID     *string            `json:"location_id"`      // 位置ID
	Status         *model.AssetStatus `json:"status"`           // 状态ID
	LineTypeID     *string            `json:"line_type_id"`     // 线路类型ID
}
