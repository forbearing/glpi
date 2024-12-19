package model

type CronTaskStatusType string

const (
	CronTaskStatusTypeScheduled CronTaskStatusType = "scheduled"
	CronTaskStatusTypeDisabled  CronTaskStatusType = "disabled"
)

type CronTaskRunMode string

const (
	CronTaskRunModeGLPI CronTaskRunMode = "glpi"
	CronTaskRunModeCLI  CronTaskRunMode = "cli"
)

type CronTaskLogStatus string

const (
	CronTaskLogStatusStart CronTaskLogStatus = "start"
	CronTaskLogStatusEnd   CronTaskLogStatus = "end"
)
