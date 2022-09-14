package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       int64  `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Lastname string `json:"lastname,omitempty"`
}
