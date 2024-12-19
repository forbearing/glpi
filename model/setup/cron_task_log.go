package model

import (
	"time"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*CronTaskLog]()
}

// CronTaskLog 计划任务日志
type CronTaskLog struct {
	pkgmodel.Base

	// 关联ID
	CronTaskID *int `json:"crontasks_id"` // 计划任务ID
	ParentID   *int `json:"parent_id"`    // 父日志ID

	// 执行信息
	Date    *time.Time          `json:"date"`    // 执行时间
	Status  *CronTaskStatusType `json:"Status"`  // 执行状态
	Elapsed *int                `json:"elapsed"` // 耗时(秒)
	Volume  *int                `json:"volume"`  // 处理数量
	Content *string             `json:"content"` // 执行内容
}
