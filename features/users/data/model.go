package data

import (
	team "clean-arch/features/teams/data"
	"clean-arch/utils/helpers"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	TeamId   uint
	Team     *team.Team `gorm:"foreignKey:TeamId"`
	FullName string
	Email    string `gorm:"unique"`
	Password string
	Role     string
	Status   bool
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.Password, err = helpers.HashPassword(user.Password)
	return
}
