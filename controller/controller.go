package controller

import "github.com/flaglogr/flaglogr/model"

type Controller struct {
	db *model.DBConnector
}

type CreatedResponse struct {
	Id uint `json:"id"`
}

func NewController(dbcon *model.DBConnector) *Controller {
	return &Controller{
		db: dbcon,
	}
}
