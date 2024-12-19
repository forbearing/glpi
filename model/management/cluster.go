package model

import (
	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*Cluster]()
}

// Cluster 集群信息
type Cluster struct {
	pkgmodel.Base

	// 基础信息
	Name    *string            `json:"name"`    // 集群名称
	UUID    *string            `json:"uuid"`    // UUID
	Version *string            `json:"version"` // 版本
	Status  *model.AssetStatus `json:"status"`  // 状态ID

	IsRecursive *bool `json:"is_recursive"` // 是否继承父级权限

	// 关联信息
	EntityID           *string `json:"entity_id"`                       // 所属实体ID
	ClusterTypeID      *string `json:"cluster_type_id"`                 // 集群类型ID
	AutoUpdateSystemID *string `json:"auto_update_system_id,omitempty"` // 自动更新系统ID
	TechUserID         *string `json:"tech_user_id"`                    // 技术用户ID
	TechGroupID        *string `json:"tech_group_id"`                   // 技术组ID
}
