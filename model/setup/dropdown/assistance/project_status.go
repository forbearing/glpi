package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*ProjectStatus]()
}

// ProjectState 项目状态
type ProjectStatus struct {
	pkgmodel.Base

	// 基础信息
	Name  *string `json:"name"`  // 状态名称
	Color *string `json:"color"` // 状态颜色

	// 状态标识
	IsFinished *bool `json:"is_finished"` // 是否已完成
}

func (*ProjectStatus) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*ProjectStatus]())
}
