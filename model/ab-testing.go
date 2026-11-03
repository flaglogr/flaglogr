package model

import "time"

type Evaluation struct {
	ID                uint `gorm:"primaryKey"`
	ProjectID         uint
	Flags             []Flag
	EvaluationResults []EvaluationResult
}

type EvaluationResult struct {
	ID           uint `gorm:"primaryKey"`
	EvaluationID uint
	SessionID    string
	Timestamp    time.Time
	FlagIndex    *uint  // The flag evaluation this belongs to. Optional
	DataType     string // Can be string, json, number
	Data         string // String encoded data
}
