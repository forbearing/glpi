package model

import "glpi/model"

type (
	TicketType    string
	TicketStatus  string
	ProblemStatus string
	PlaningStatus string
	ChangeStatus  string
	Impact        string
	Urgency       string
	Priority      string
)

const (
	TicketTypeIncident TicketType = "incident"
	TicketTypeRequest  TicketType = "request"

	TicketStatusNew                TicketStatus = model.STATUS_NEW
	TicketStatusProcessingAssigned TicketStatus = model.STATUS_PROCESSING_ASSIGNED
	TicketStatusProcessingPlanned  TicketStatus = model.STATUS_PROCESSING_PLANNED
	TicketStatusPending            TicketStatus = model.STATUS_PENDING
	TicketStatusSolved             TicketStatus = model.STATUS_SOLVED
	TicketStatusClosed             TicketStatus = model.STATUS_CLOSED

	ProblemStatusNew                ProblemStatus = model.STATUS_NEW
	ProblemStatusAccepted           ProblemStatus = model.STATUS_ACCEPTED
	ProblemStatusProcessingAssigned ProblemStatus = model.STATUS_PROCESSING_ASSIGNED
	ProblemStatusProcessingPlanned  ProblemStatus = model.STATUS_PROCESSING_PLANNED
	ProblemStatusPending            ProblemStatus = model.STATUS_PENDING
	ProblemStatusSolved             ProblemStatus = model.STATUS_SOLVED
	ProblemStatusUnderObservation   ProblemStatus = model.STATUS_UNDER_OBSERVATION
	ProblemStatusClosed             ProblemStatus = model.STATUS_CLOSED

	ChangeStatusNew           ChangeStatus = model.STATUS_NEW
	ChangeStatusEvaluation    ChangeStatus = model.STATUS_EVALUATION
	ChangeStatusApproval      ChangeStatus = model.STATUS_APPROVAL
	ChangeStatusAccepted      ChangeStatus = model.STATUS_ACCEPTED
	ChangeStatusPending       ChangeStatus = model.STATUS_PENDING
	ChangeStatusTesting       ChangeStatus = model.STATUS_TESTING
	ChangeStatusQualification ChangeStatus = model.STATUS_QUALIFICATION
	ChangeStatusApplied       ChangeStatus = model.STATUS_APPLIED
	ChangeStatusReview        ChangeStatus = model.STATUS_REVIEW
	ChangeStatusClosed        ChangeStatus = model.STATUS_CLOSED
	ChangeStatusCancelled     ChangeStatus = model.STATUS_CANCELLED
	ChangeStatusRefused       ChangeStatus = model.STATUS_REFUSED

	PlaningStatusInformation PlaningStatus = model.STATUS_INFORMATION
	PlaningStatusTodo        PlaningStatus = model.STATUS_TODO
	PlaningStatusDone        PlaningStatus = model.STATUS_DONE

	ImpactMajor    Impact = model.MAJOR
	ImpactVeryHigh Impact = model.VERY_HIGH
	ImpactHigh     Impact = model.HIGH
	ImpactMedium   Impact = model.MEDIUM
	ImpactLow      Impact = model.LOW
	ImpactVeryLow  Impact = model.VERY_LOW

	UrgencyVeryHigh Urgency = model.VERY_HIGH
	UrgencyHigh     Urgency = model.HIGH
	UrgencyMedium   Urgency = model.MEDIUM
	UrgencyLow      Urgency = model.LOW
	UrgencyVeryLow  Urgency = model.VERY_LOW

	PriorityMajor    Priority = model.MAJOR
	PriorityVeryHigh Priority = model.VERY_HIGH
	PriorityHigh     Priority = model.HIGH
	PriorityMedium   Priority = model.MEDIUM
	PriorityLow      Priority = model.LOW
	PriorityVeryLow  Priority = model.VERY_LOW
)
