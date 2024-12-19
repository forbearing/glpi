package model

import (
	"fmt"

	"glpi/model"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*DeviceCamera]()
}

// DeviceCamera 设备摄像头
type DeviceCamera struct {
	pkgmodel.Base

	Designation *string `json:"designation"` // 摄像头标识

	// 组织相关
	EntityID    *string `json:"entity_id"`    // 所属实体ID
	IsRecursive *bool   `json:"is_recursive"` // 是否递归继承

	// 制造商信息
	ManufacturerID *string `json:"manufacturer_id"` // 制造商ID
	ModelID        *string `json:"model_id"`        // 型号ID

	// 摄像头参数
	IsFlashUnit *bool   `json:"is_flashunit"` // 是否有闪光灯
	LensFacing  *string `json:"lensfacing"`   // 镜头朝向
	Orientation *string `json:"orientation"`  // 安装方向
	FocalLength *string `json:"focal_length"` // 焦距
	SensorSize  *string `json:"sensor_size"`  // 传感器尺寸
	Support     *string `json:"support"`      // 支持特性
}

func (*DeviceCamera) GetTableName() string {
	return fmt.Sprintf("%s_%s", model.PREFIX_COMPONENT, pkgmodel.GetTableName[*DeviceCamera]())
}
