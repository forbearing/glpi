package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
	"github.com/forbearing/golib/util"
)

var (
	ValueOf = util.ValueOf[string]
	Deref   = util.Deref[string]
)

var variables = []*SSOVariable{
	{model.Name{Name: ValueOf("HTTP_AUTH_USER")}},
	{model.Name{Name: ValueOf("HTTP_REMOTE_USER")}},
	{model.Name{Name: ValueOf("PHP_AUTH_USER")}},
	{model.Name{Name: ValueOf("REDIRECT_REMOTE_USER")}},
	{model.Name{Name: ValueOf("REMOTE_USER")}},
	{model.Name{Name: ValueOf("USERNAME")}},
}

func init() {
	for i := range variables {
		variables[i].ID = Deref(variables[i].Name.Name)
	}
	pkgmodel.Register(variables...)
}

type SSOVariable struct {
	model.Name
}

func (*SSOVariable) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*SSOVariable]())
}
