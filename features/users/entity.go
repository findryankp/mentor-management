package users

import (
	"immersiveApp/features/teams"
	"time"
)

type UserEntity struct {
	Id          uint
	TeamId      uint   `validate:"required"`
	FullName    string `validate:"required"`
	Email       string `validate:"required,email"`
	Password    string `validate:"required"`
	PhoneNumber string `validate:"required"`
	Address     string `validate:"required"`
	Role        string `validate:"required"`
	Team        teams.TeamEntity
	Status      bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UserServiceInterface interface {
	GetAll() ([]UserEntity, error)
	GetById(id uint) (UserEntity, error)
	Create(userEntity UserEntity) (UserEntity, error)
	Update(userEntity UserEntity, id uint) (UserEntity, error)
	Delete(id uint) error
}

type UserDataInterface interface {
	SelectAll() ([]UserEntity, error)
	SelectById(id uint) (UserEntity, error)
	Store(userEntity UserEntity) (uint, error)
	Edit(userEntity UserEntity, id uint) (uint, error)
	Destroy(id uint) error
}
