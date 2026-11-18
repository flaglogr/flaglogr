package model

type Evaluation struct {
	ID                uint               `gorm:"primaryKey"`
	ProjectID         uint               `gorm:"-"`
	Flags             []Flag             `gorm:"constraint:OnDelete:CASCADE;"`
	EvaluationResults []EvaluationResult `gorm:"constraint:OnDelete:CASCADE;"`
}

type EvaluationResult struct {
	ID           uint `gorm:"primaryKey"`
	EvaluationID uint
	FlagIndex    *uint // The flag evaluation this belongs to. Optional
	AnalyticsID  uint  `gorm:"constraint:OnDelete:CASCADE;"`
}
