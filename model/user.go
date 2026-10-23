package model

type User struct {
	ID       uint `gorm:"primaryKey"`
	OidcId   string
	Username string
	Groups   []Group `gorm:"many2many:user_groups;"`
}

type Group struct {
	ID       uint `gorm:"primaryKey"`
	Projects []Project
	IsAdmin  bool
	Name     string
}
