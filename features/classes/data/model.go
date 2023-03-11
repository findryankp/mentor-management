package data

import (
	data "immersiveApp/features/users/data"

	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	ClassName    string
	PicUserId    uint
	User         *data.User `gorm:"foreignKey:PicUserId"`
	DateStart    string
	DateGraduate string
}
