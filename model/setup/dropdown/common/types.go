package model

type BlacklistType string

const (
	BlacklistTypeIP           BlacklistType = "ip"
	BlacklistTypeMAC          BlacklistType = "mac"
	BlacklistTypeSerialNumber BlacklistType = "serial_number"
	BlacklistTypeUUID         BlacklistType = "uuid"
	BlacklistTypeEmail        BlacklistType = "email"
	BlacklistTypeModel        BlacklistType = "model"
)
