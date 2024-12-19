package model

import pkgmodel "github.com/forbearing/golib/model"

func init() {
	pkgmodel.Register[*Document]()
}

// Document 文档信息
type Document struct {
	pkgmodel.Base

	// 基础信息
	Name        *string `json:"name"`         // 文档名称
	IsRecursive *bool   `json:"is_recursive"` // 是否继承父级权限
	Tag         *string `json:"tag"`          // 文档标签

	// 文件信息
	Filename *string `json:"filename"` // 文件名
	Filepath *string `json:"filepath"` // 文件路径
	Mime     *string `json:"mime"`     // MIME类型
	Sha1sum  *string `json:"sha1sum"`  // SHA1校验和
	Link     *string `json:"link"`     // 外部链接

	// 状态信息
	IsBlacklisted *bool `json:"is_blacklisted"` // 是否被拉黑

	// 关联信息
	EntityID           *string `json:"entity_id"`            // 所属实体ID
	DocumentCategoryID *string `json:"document_category_id"` // 文档分类ID
	UserID             *string `json:"user_id"`              // 用户ID
	TicketID           *string `json:"ticket_id"`            // 工单ID
}
