package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*RuleRightParameter]()
}

// RuleRightParameter 规则权限参数
type RuleRightParameter struct {
	pkgmodel.Base

	Name  *string `json:"name"`  // 参数名称
	Value *string `json:"value"` // 参数值
}

func (*RuleRightParameter) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*RuleRightParameter]())
}
