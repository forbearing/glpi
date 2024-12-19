package model

import "glpi/model"

type (
	ProjectStatus   string
	ReminderStatus  string
	ProjectPriority string
)

const (
	ProjectStatusClosed     ProjectStatus = model.STATUS_CLOSED
	ProjectStatusNew        ProjectStatus = model.STATUS_NEW
	ProjectStatusProcessing ProjectStatus = model.STATUS_PROCESSING

	ReminderStatusInformation ReminderStatus = model.STATUS_INFORMATION
	ReminderStatusTODO        ReminderStatus = model.STATUS_TODO
	ReminderStatusDONE        ReminderStatus = model.STATUS_DONE

	ProjectPriorityMajor    ProjectPriority = model.MAJOR
	ProjectPriorityVeryHigh ProjectPriority = model.VERY_HIGH
	ProjectPriorityHigh     ProjectPriority = model.HIGH
	ProjectPriorityMedium   ProjectPriority = model.MEDIUM
	ProjectPriorityLow      ProjectPriority = model.LOW
	ProjectPriorityVeryLow  ProjectPriority = model.VERY_LOW
)
