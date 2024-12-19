package model

import (
	"time"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*Database]()
}

// Database 数据库信息
type Database struct {
	pkgmodel.Base

	// 基础信息
	Name     *string    `json:"name"`      // 数据库名称
	Size     *int64     `json:"size"`      // 数据库大小
	IsActive *bool      `json:"is_active"` // 是否激活
	LastSync *time.Time `json:"last_sync"` // 最后同步时间

	// 备份相关
	IsOnBackup   *bool      `json:"is_on_backup"`   // 是否在备份中
	LastBackupAt *time.Time `json:"last_backup_at"` // 最后备份时间

	// 系统相关
	IsRecursive *bool `json:"is_recursive"` // 是否继承父级权限
	IsDynamic   *bool `json:"is_dynamic"`   // 是否动态

	// 关联信息
	EntityID           *string `json:"entity_id"`            // 所属实体ID
	DatabaseInstanceID *string `json:"database_instance_id"` // 数据库实例ID
}
