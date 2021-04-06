package models

import (
	"github.com/jinzhu/gorm"
)

type Note struct {
	gorm.Model
	Title string `gorm:"type:varchar(255);" json:"title"`
	Body  string `gorm:"type:text" json:"body"`
}

func GetNote() Note {
	var note Note
	return note
}

func GetNotes() []Note {
	var notes []Note
	return notes
}
