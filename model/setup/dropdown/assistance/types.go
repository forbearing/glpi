package model

import "glpi/model"

type (
	TaskTemplateType                  string
	PlanningExternalEventTemplateType string
)

const (
	TaskTemplateTypeInformation TaskTemplateType = model.STATUS_INFORMATION
	TaskTemplateTypeTodo        TaskTemplateType = model.STATUS_TODO
	TaskTemplateTypeDone        TaskTemplateType = model.STATUS_DONE

	PlanningExternalEventTemplateTypeInformation PlanningExternalEventTemplateType = model.STATUS_INFORMATION
	PlanningExternalEventTemplateTypeTodo        PlanningExternalEventTemplateType = model.STATUS_TODO
	PlanningExternalEventTemplateTypeDone        PlanningExternalEventTemplateType = model.STATUS_DONE
)
