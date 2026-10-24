package model

type Project struct {
	ID        uint `gorm:"primaryKey"`
	APIKeys []APIKey
}

type APIKey struct {
	ID        uint `gorm:"primaryKey"`
	ProjectID uint
	Name      string
	Key       string
}
