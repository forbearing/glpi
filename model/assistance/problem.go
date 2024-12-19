package model

import (
	"time"

	pkgmodel "github.com/forbearing/golib/model"
)

func init() {
	pkgmodel.Register[*Problem]()
}

// Problem 问题记录
type Problem struct {
	pkgmodel.Base

	// 基础信息
	Name    *string        `json:"name"`    // 问题标题
	Content *string        `json:"content"` // 问题描述
	Status  *ProblemStatus `json:"status"`  // 问题状态

	// 时间信息
	OpenDate         *time.Time `json:"open_date"`          // 创建时间
	CloseDate        *time.Time `json:"close_date"`         // 关闭时间
	SolveDate        *time.Time `json:"solve_date"`         // 解决时间
	TimeToResolve    *time.Time `json:"time_to_resolve"`    // 解决期限
	BeginWaitingDate *time.Time `json:"begin_waiting_date"` // 等待开始时间

	// 影响评估
	Urgency       *Urgency  `json:"urgency"`        // 紧急程度
	Impact        *Impact   `json:"impact"`         // 影响程度
	Priority      *Priority `json:"priority"`       // 优先级别
	ImpactContent *string   `json:"impact_content"` // 影响内容描述

	// 问题分析
	CauseContent   *string `json:"cause_content"`   // 原因分析
	SymptomContent *string `json:"symptom_content"` // 症状描述

	// 组织信息
	EntityID    *string `json:"entity_id"`    // 所属实体ID
	LocationID  *string `json:"location_id"`  // 位置ID
	CategoryID  *string `json:"category_id"`  // 问题类别ID
	IsRecursive *bool   `json:"is_recursive"` // 是否继承父级权限

	// 用户信息
	LastUpdaterID *string `json:"last_updater_id"` // 最后更新人ID
	RecipientID   *string `json:"recipient_id"`    // 接收人ID

	// 统计信息
	ActionTime      *int64 `json:"action_time"`      // 操作时长
	WaitingDuration *int64 `json:"waiting_duration"` // 等待时长
	CloseDelayStat  *int64 `json:"close_delay_stat"` // 关闭延迟统计
	SolveDelayStat  *int64 `json:"solve_delay_stat"` // 解决延迟统计
}
