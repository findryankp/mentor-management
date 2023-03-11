package classes

import (
	"immersiveApp/features/users"
	"time"
)

type ClassEntity struct {
	Id           uint
	ClassName    string `validate:"required"`
	PicUserId    uint   `validate:"required"`
	User         users.UserEntity
	DateStart    string `validate:"required"`
	DateGraduate string `validate:"required"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type ClassServiceInterface interface {
	GetAll() ([]ClassEntity, error)
	GetById(id uint) (ClassEntity, error)
	Create(classEntity ClassEntity) (ClassEntity, error)
	Update(classEntity ClassEntity, id uint) (ClassEntity, error)
	Delete(id uint) error
}

type ClassDataInterface interface {
	SelectAll() ([]ClassEntity, error)
	SelectById(id uint) (ClassEntity, error)
	Store(classEntity ClassEntity) (uint, error)
	Edit(classEntity ClassEntity, id uint) error
	Destroy(id uint) error
}
