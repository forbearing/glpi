package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*PlanningEventCategory]()
}

// PlanningEventCategory 计划事件分类
type PlanningEventCategory struct {
	pkgmodel.Base

	// 基础信息
	Name  *string `json:"name"`  // 分类名称
	Color *string `json:"color"` // 分类颜色
}

func (*PlanningEventCategory) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*PlanningEventCategory]())
}
