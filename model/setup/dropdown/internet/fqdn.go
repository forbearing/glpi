package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*FQDN]()
}

// FQDN 完全限定域名
type FQDN struct {
	model.NameEntity

	// 域名配置
	FQDN *string `json:"fqdn"` // 完全限定域名
}

func (*FQDN) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*FQDN]())
}
