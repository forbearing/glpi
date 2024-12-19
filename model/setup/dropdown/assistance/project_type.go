package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*ProjectType]()
}

// ProjectType 项目类型
type ProjectType struct {
	pkgmodel.Base

	// 基础信息
	Name *string `json:"name"` // 类型名称
}

func (*ProjectType) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*ProjectType]())
}
