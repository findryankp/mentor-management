package service

import (
	"clean-arch/app/middlewares"
	"clean-arch/features/auth"
	"clean-arch/features/users"
	"clean-arch/utils/helpers"
	"errors"
)

type authService struct {
	data auth.AuthDataInterface
}

func New(u auth.AuthDataInterface) auth.AuthServiceInterface {
	return &authService{
		data: u,
	}
}

func (u *authService) Login(email, password string) (string, error) {
	if email == "" || password == "" {
		return "", errors.New("email and password must be fill")
	}

	user, err := u.data.GetUserByEmail(email)
	if err != nil || !helpers.CheckPasswordHash(password, user.Password) {
		return "", errors.New("user and password not found")
	}

	//make jwt
	token, errToken := middlewares.CreateToken(int(user.Id), user.Role)
	if errToken != nil {
		return "", errToken
	}

	return token, nil
}

func (u *authService) Register(request users.UserEntity) error {
	return u.data.Register(request)
}
