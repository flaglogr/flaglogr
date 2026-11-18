package model

type User struct {
	ID       uint `gorm:"primaryKey"`
	OidcId   string
	Username string
	IsAdmin  bool
	Projects []Project `gorm:"many2many:user_projects"`
}
