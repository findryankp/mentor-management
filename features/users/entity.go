package users

import "time"

type UserEntity struct {
	Id        uint
	Name      string
	Email     string
	Password  string
	Address   string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserServiceInterface interface {
	GetAll() ([]UserEntity, error)
	GetById(id uint) (UserEntity, error)
	Create(userEntity UserEntity) error
	Update(userEntity UserEntity, id uint) error
	Delete(id uint) error
}

type UserDataInterface interface {
	SelectAll() ([]UserEntity, error)
	SelectById(id uint) (UserEntity, error)
	Store(userEntity UserEntity) error
	Edit(userEntity UserEntity, id uint) error
	Destroy(id uint) error
}
