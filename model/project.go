package model

type Project struct {
	ID          uint         `json:"-" gorm:"primaryKey"`
	Name        string       `json:"name"`
	APIKeys     []APIKey     `json:"apiKeys" gorm:"constraint:OnDelete:CASCADE;"`
	Evaluations []Evaluation `json:"-" gorm:"constraint:OnDelete:CASCADE;"`
	Analytics   []Analytics  `json:"-" gorm:"constraint:OnDelete:CASCADE"`
}

type APIKey struct {
	ID        uint   `json:"-" gorm:"primaryKey"`
	ProjectID uint   `json:"-"`
	Name      string `json:"name"`
	Key       string `json:"-"`
}
