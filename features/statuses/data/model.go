package data

import "gorm.io/gorm"

type Status struct {
	gorm.Model
	Name     string
	ParentId string
	Desc     string
}
