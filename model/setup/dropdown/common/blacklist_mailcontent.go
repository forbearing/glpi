package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*BlacklistedMailContent]()
}

// BlacklistedMailContent 邮件黑名单内容
type BlacklistedMailContent struct {
	pkgmodel.Base

	// 基础信息
	Name    *string `json:"name"`    // 黑名单名称
	Content *string `json:"content"` // 黑名单内容
}

func (*BlacklistedMailContent) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*BlacklistedMailContent]())
}
