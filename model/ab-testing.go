package model

import "time"

type Evaluation struct {
	ID                uint `gorm:"primaryKey"`
	ProjectID         uint
	Flags             []Flag             `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	EvaluationResults []EvaluationResult `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type EvaluationResult struct {
	ID           uint `gorm:"primaryKey"`
	EvaluationID uint
	FlagIndex    *uint // The flag evaluation this belongs to. Optional
	AnalyticsID  uint  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Analytics struct {
	ID        uint `gorm:"primaryKey"`
	SessionID string
	Timestamp time.Time
	DataType  string // Can be string, json, number
	Data      string // String encoded data
	Event     string // name of the event that was recorded
}
