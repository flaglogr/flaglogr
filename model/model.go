package model

import (
	"context"
	"fmt"

	"github.com/TwiN/logr"
	"github.com/flaglogr/flaglogr/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBConnector struct {
	DB     *gorm.DB
	Config *config.Config
}

var ctx = context.Background()

func Init(prod bool, conf *config.Config) *DBConnector {
	var gormConf gorm.Config
	if prod {
		gormConf = gorm.Config{}
	} else {
		gormConf = gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		}
	}

	db, err := gorm.Open(sqlite.Open("flaglogr.db"), &gormConf)
	if err != nil {
		logr.Fatal(fmt.Sprint("[DB] Failed to connect to DB with error", err))
	}

	// FIXME: Only used if using sqlite db
	db.Exec("PRAGMA foreign_keys = ON")

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Evaluation{})
	db.AutoMigrate(&EvaluationResult{})
	db.AutoMigrate(&Flag{})
	db.AutoMigrate(&FlagValues{})
	db.AutoMigrate(&Analytics{})
	db.AutoMigrate(&Project{})
	db.AutoMigrate(&APIKey{})

	return &DBConnector{
		DB:     db,
		Config: conf,
	}
}

func dereferenceOrDefault[T any](val *T, defaultValue T) T {
	if val != nil {
		return *val
	} else {
		return defaultValue
	}
}
