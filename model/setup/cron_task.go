package model

import (
	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*CronTask]()
}

// CronTask 计划任务
type CronTask struct {
	pkgmodel.Base

	Name *string `json:"name"` // 任务名称

	// 任务类型
	ItemType *string `json:"item_type"` // 对象类型

	// 执行配置
	Frequency *int                `json:"frequency"`  // 执行频率(秒)
	Param     *int                `json:"param"`      // 参数
	Status    *CronTaskStatusType `json:"status"`     // 状态
	RunMode   *CronTaskRunMode    `json:"run_mode"`   // 执行模式
	AllowMode *int                `json:"allow_mode"` // 允许模式

	// 时间窗口
	PeriodFrom *int `json:"period_from"` // 最小小时
	PeriodTo   *int `json:"period_to"`   // 最大小时

	// 日志相关
	LogLifetime *int    `json:"log_lifetime"` // 日志保留天数
	LastRun     *string `json:"last_run"`     // 上次运行时间
	LastCode    *string `json:"last_code"`    // 上次执行代码

	Description *string `json:"description"`
}
