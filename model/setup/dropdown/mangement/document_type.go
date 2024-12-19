package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*DocumentType]()
}

// DocumentType 文档类型
type DocumentType struct {
	pkgmodel.Base

	Name         *string `json:"name"`          // 名称
	Extension    *string `json:"extension"`     // 文件扩展名
	Icon         *string `json:"icon"`          // 图标文件名
	MIME         *string `json:"mime"`          // MIME类型
	IsUploadable *bool   `json:"is_uploadable"` // 是否允许上传
}

func (*DocumentType) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_DROPDOWN, pkgmodel.GetTableName[*DocumentType]())
}
