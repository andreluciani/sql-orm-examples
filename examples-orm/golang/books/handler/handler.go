package handler

import "gorm.io/gorm"

func NewController(db *gorm.DB) *Controller {
	return &Controller{
		db: *db,
	}
}

type Controller struct {
	db gorm.DB
}
