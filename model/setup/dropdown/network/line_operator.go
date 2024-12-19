package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*LineOperator]()
}

// LineOperator 线路运营商
type LineOperator struct {
	pkgmodel.Base

	Name *string `json:"name"` // 名称

	// 组织相关
	EntityID    *string `json:"entity_id"`    // 所属实体ID
	IsRecursive *bool   `json:"is_recursive"` // 是否递归继承

	// 移动运营商标识
	MCC *int `json:"mcc"` // 移动国家代码(Mobile Country Code)
	MNC *int `json:"mnc"` // 移动网络代码(Mobile Network Code)
}

func (*LineOperator) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*LineOperator]())
}
