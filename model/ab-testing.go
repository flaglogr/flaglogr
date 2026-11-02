package model

import "time"

type Evaluation struct {
	ID    uint `gorm:"primaryKey"`
	Flags []Flag
}

type EvaluationResults struct {
	ID           uint `gorm:"primaryKey"`
	EvaluationID uint
	SessionID    string
	Timestamp    time.Time
	DataType     string
	Data         string // String encoded data
}

type EvaluationSession struct{}
