package model

type User struct {
	ID       uint `gorm:"primaryKey"`
	OidcId   string
	Username string
	Groups   []Group `gorm:"many2many:user_groups;constraint:OnDelete:CASCADE"`
}

// A group is what manages projects and has a number of users
type Group struct {
	ID       uint `gorm:"primaryKey"`
	Projects []Project
	IsAdmin  bool
	Name     string
	Users    []User `gorm:"many2many:user_groups;constraint:OnDelete:CASCADE"`
}
