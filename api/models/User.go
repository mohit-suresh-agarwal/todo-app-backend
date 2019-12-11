package models

import(
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name  string    `gorm:"size:255;not null;unique" json:"name"`
	Email     string    `gorm:"size:100;not null;unique" json:"email"`
	Password  string    `gorm:"size:100;not null;" json:"password"`
	Tasks    	[]Task		`json:"tasks"`
}