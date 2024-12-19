package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*ProjectTaskType]()
}

// ProjectTaskType 项目类型
type ProjectTaskType struct {
	pkgmodel.Base

	// 基础信息
	Name *string `json:"name"` // 类型名称
}

func (*ProjectTaskType) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*ProjectTaskType]())
}
