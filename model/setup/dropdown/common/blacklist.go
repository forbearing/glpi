package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*Blacklist]()
}

// Blacklist 黑名单信息
type Blacklist struct {
	pkgmodel.Base

	// 基础信息
	Name  *string        `json:"name"`  // 黑名单名称
	Type  *BlacklistType `json:"type"`  // 黑名单类型
	Value *string        `json:"value"` // 黑名单值
}

func (*Blacklist) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*Blacklist]())
}
