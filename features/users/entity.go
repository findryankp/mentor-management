package users

import (
	"time"
)

type UserEntity struct {
	Id        uint
	TeamId    uint
	FullName  string
	Email     string
	Password  string
	Role      string
	Status    bool
	CreatedAt time.Time
	UpdatedAt time.Time
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
	Store(userEntity UserEntity) (UserEntity, error)
	Edit(userEntity UserEntity, id uint) (UserEntity, error)
	Destroy(id uint) error
}
