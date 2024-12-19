package model

type AssetStatus string

const (
	AssetStatusUsed        AssetStatus = "used"        // 在用
	AssetStatusInventory   AssetStatus = "inventory"   // 库存
	AssetStatusLended      AssetStatus = "lended"      // 借出
	AssetStatusAbnormal    AssetStatus = "abnormal"    // 异常
	AssetStatusMaintenance AssetStatus = "maintenance" // 维修
	AssetStatusRetired     AssetStatus = "retired"     // 报废
)

const (
	STATUS_NEW                 = "new"
	STATUS_EVALUATION          = "evaluation"
	STATUS_APPROVAL            = "approval"
	STATUS_ACCEPTED            = "accepted"
	STATUS_PROCESSING          = "processing"
	STATUS_PROCESSING_ASSIGNED = "processing_assigned"
	STATUS_PROCESSING_PLANNED  = "processing_planned"
	STATUS_PENDING             = "pending"
	STATUS_TESTING             = "testing"
	STATUS_SOLVED              = "solved"
	STATUS_UNDER_OBSERVATION   = "under_observation"
	STATUS_QUALIFICATION       = "qualification"
	STATUS_APPLIED             = "applied"
	STATUS_REVIEW              = "review"
	STATUS_CLOSED              = "closed"
	STATUS_CANCELLED           = "cancelled"
	STATUS_REFUSED             = "refused"

	STATUS_INFORMATION = "information"
	STATUS_TODO        = "todo"
	STATUS_DONE        = "done"

	MAJOR     = "major"
	VERY_HIGH = "very_high"
	HIGH      = "high"
	MEDIUM    = "medium"
	LOW       = "low"
	VERY_LOW  = "very_low"
)

const (
	PREFIX_DROPDOWN  = "dropdown"
	PREFIX_COMPONENT = "component"
)
