package model

type Flag struct {
	ID          uint   `gorm:"primaryKey"`
	Identifier  string // The ID used by the application
	Description string // Any extra description in for the flag (up to 250 chars)
	ValueType   string // Can be `number` | `bool` | `json` | `string`
	Value       string // Value encoded as string for easier storage
}
