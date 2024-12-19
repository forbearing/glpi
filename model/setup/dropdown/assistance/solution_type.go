package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*SolutionType]()
}

// SolutionType 解决方案类型
type SolutionType struct {
	pkgmodel.Base

	// 基础信息
	Name *string `json:"name"` // 类型名称

	// 组织信息
	EntityID    *string `json:"entity_id"`    // 实体ID
	IsRecursive *bool   `json:"is_recursive"` // 是否递归
}

func (*SolutionType) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*SolutionType]())
}
