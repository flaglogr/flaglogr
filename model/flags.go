package model

type Flag struct {
	ID           uint         `json:"-" gorm:"primaryKey"`
	Identifier   string       `json:"identifier"`  // The ID used by the application
	Description  string       `json:"description"` // Any extra description in for the flag (up to 250 chars)
	Values       []FlagValues `json:"values"`
	EvaluationID *uint        `json:"-"`
	Enabled      bool         `json:"enabled"`
}

type FlagValues struct {
	ID          uint   `json:"valueId" gorm:"primaryKey"` // The ID of the value that is used to identify which one was selected
	FlagID      uint   `json:"-"`
	DeployShare uint   `json:"deployShare"` // Can be between 0 and 100 (i.e. percentage, must be integer) if just one, for the flag, then ignored
	ValueType   string `json:"valueType"`
	Value       string `json:"value"`
}
