package model

import "gorm.io/gorm"

type DBConnector struct {
	DB *gorm.DB
}
