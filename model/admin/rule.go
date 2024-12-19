package model

import pkgmodel "github.com/forbearing/golib/model"

func init() {
	pkgmodel.Register[*Rule]()
}

// Rule 规则信息
type Rule struct {
	pkgmodel.Base

	// 基础信息
	Name        *string `json:"name"`        // 规则名称
	Description *string `json:"description"` // 规则描述
	UUID        *string `json:"uuid"`        // 唯一标识符
	IsActive    *bool   `json:"is_active"`   // 是否启用

	// 组织信息
	EntityID    *string `json:"entity_id"`    // 实体ID
	IsRecursive *bool   `json:"is_recursive"` // 是否递归应用到子实体

	// 规则配置
	SubType   *string `json:"sub_type"`  // 规则子类型(如: RuleImportAsset)
	Ranking   *int    `json:"ranking"`   // 规则执行顺序
	Match     *string `json:"match"`     // 条件匹配方式(AND/OR)
	Condition *int    `json:"condition"` // 条件类型
}
