package model

type Flag struct {
	ID          uint   `gorm:"primaryKey"`
	Identifier  string // The ID used by the application
	Description string // Any extra description in for the flag (up to 250 chars)
	Values      []FlagValues
	Enabled     bool
}

type FlagValues struct {
	FlagID      uint
	DeployShare uint // Can be between 0 and 100 (i.e. percentage, must be integer)
	ID          uint `gorm:"primaryKey"`
	ValueType   string
	Value       string
}
