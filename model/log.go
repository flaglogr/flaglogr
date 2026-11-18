package model

import "time"

type Analytics struct {
	ID        uint      `gorm:"primaryKey" json:"-"`
	ProjectID uint      `json:"-"`         // The project this belongs to
	SessionID string    `json:"sessionId"` // ID of the session cookie
	Timestamp time.Time `json:"timestamp"` // Time of recording
	DataType  string    `json:"dataType"`  // Can be string, json, number
	Data      string    `json:"data"`      // String encoded data
	Event     string    `json:"event"`     // name of the event that was recorded
}
