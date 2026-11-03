package model

type Project struct {
	ID          uint         `gorm:"primaryKey"`
	APIKeys     []APIKey     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Evaluations []Evaluation `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type APIKey struct {
	ID        uint `gorm:"primaryKey"`
	ProjectID uint
	Name      string
	Key       string
}
