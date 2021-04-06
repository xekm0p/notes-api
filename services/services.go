package services

import (
	"restapi/models"

	"github.com/jinzhu/gorm"
)

var dbconn *gorm.DB

type Response struct {
	Data    []models.Note `json:"data"`
	Message string        `json:"message"`
}

func SetDB(db *gorm.DB) {
	dbconn = db
	var note = models.GetNote()
	dbconn.AutoMigrate(&note)
}
