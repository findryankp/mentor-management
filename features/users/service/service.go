package service

import "clean-arch/features/users"

type userService struct {
	Data users.UserDataInterface
}

func New(data users.UserDataInterface) users.UserServiceInterface {
	return &userService{
		Data: data,
	}
}

func (u *userService) GetAll() ([]users.UserEntity, error) {
	return u.Data.SelectAll()
}

func (u *userService) GetById(id uint) (users.UserEntity, error) {
	return u.Data.SelectById(id)
}

func (u *userService) Create(request users.UserEntity) error {
	return u.Data.Store(request)
}

func (u *userService) Update(request users.UserEntity, id uint) error {
	return u.Data.Edit(request, id)
}

func (u *userService) Delete(id uint) error {
	return u.Data.Destroy(id)
}
