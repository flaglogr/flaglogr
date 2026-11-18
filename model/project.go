package model

import "time"

type Project struct {
	ID          uint         `json:"-" gorm:"primaryKey"`
	UserID      uint         `json:"-"`
	GroupID     *uint        `json:"-"`
	Name        string       `json:"name"`
	UpdatedAt   time.Time    `json:"updatedAt"`
	APIKeys     []APIKey     `json:"apiKeys" gorm:"constraint:OnDelete:CASCADE;"`
	Evaluations []Evaluation `json:"-" gorm:"constraint:OnDelete:CASCADE;"`
	Analytics   []Analytics  `json:"-" gorm:"constraint:OnDelete:CASCADE;"`
}

type APIKey struct {
	ID        uint   `json:"-" gorm:"primaryKey"`
	ProjectID uint   `json:"-"`
	Name      string `json:"name"`
	Key       string `json:"-"`
}

type CreateProjectRequest struct {
	Name string `json:"name" example:"name" maxLength:"100" minLength:"3" validate:"required" binding:"required,min=3,max=100,printascii"`
}

type UpdateProjectRequest struct {
	Name    *string `json:"name" example:"name" maxLength:"250" minLength:"3" validate:"optional" binding:"omitempty,min=3,max=250,printascii"`
	GroupID *uint   `json:"groupId" example:"1" validate:"optional" binding:"omitempty,numeric"`
}

type CreateAPIKeyRequest struct {
	Name string `json:"name" example:"name" maxLength:"100" minLength:"3" validate:"required" binding:"required,min=3,max=100,printascii"`
}
