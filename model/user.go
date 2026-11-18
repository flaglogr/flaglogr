package model

type User struct {
	ID       uint `gorm:"primaryKey"`
	OidcId   string
	Username string
	IsAdmin  bool
	Projects []Project
	Groups   []Group `gorm:"many2many:user_groups;constraint:OnDelete:CASCADE"`
}

// A group is what manages projects and has a number of users
type Group struct {
	ID       uint `gorm:"primaryKey"`
	Projects []Project
	Name     string
	Users    []User `gorm:"many2many:user_groups;constraint:OnDelete:CASCADE"`
}
