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

func (s *userService) GetAll() ([]users.UserEntity, error) {
	return s.Data.SelectAll()
}

func (s *userService) GetById(id uint) (users.UserEntity, error) {
	return s.Data.SelectById(id)
}

func (s *userService) Create(request users.UserEntity) (users.UserEntity, error) {
	request.Status = true
	user_id, err := s.Data.Store(request)
	if err != nil {
		return users.UserEntity{}, err
	}

	return s.Data.SelectById(user_id)
}

func (s *userService) Update(request users.UserEntity, id uint) (users.UserEntity, error) {
	if checkDataExist, err := s.Data.SelectById(id); err != nil {
		return checkDataExist, err
	}

	user_id, err := s.Data.Edit(request, id)
	if err != nil {
		return users.UserEntity{}, err
	}

	return s.Data.SelectById(user_id)
}

func (s *userService) Delete(id uint) error {
	if _, err := s.Data.SelectById(id); err != nil {
		return err
	}

	return s.Data.Destroy(id)
}
