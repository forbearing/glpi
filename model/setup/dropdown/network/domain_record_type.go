package model

import (
	"encoding/json"
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*DomainRecordType]()
}

type DomainRecordType struct {
	model.NameEntity
	Fields json.RawMessage
}

func (*DomainRecordType) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*DomainRecordType]())
}
