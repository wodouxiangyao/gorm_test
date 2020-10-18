package model

import (
	"database/sql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name sql.NullString `gorm:"type:string;default:'\xF0\x9F\x98\x88'"`
	Age uint16
}
