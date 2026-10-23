package model

type Project struct {
}

type APIKey struct {
	ID   uint `gorm:"primaryKey"`
	Name string
	Key  string
}
