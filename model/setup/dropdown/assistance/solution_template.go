package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*SolutionTemplate]()
}

// SolutionTemplate 解决方案模板
type SolutionTemplate struct {
	pkgmodel.Base

	// 基础信息
	Name    *string `json:"name"`    // 模板名称
	Content *string `json:"content"` // 模板内容

	// 组织信息
	EntityID    *string `json:"entity_id"`    // 实体ID
	IsRecursive *bool   `json:"is_recursive"` // 是否递归

	// 类型关联
	SolutionTypeID *int `json:"solution_type_id"` // 解决方案类型ID
}

func (*SolutionTemplate) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*SolutionTemplate]())
}
