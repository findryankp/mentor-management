package auth

import (
	"clean-arch/features/users"
)

type AuthServiceInterface interface {
	Login(email, password string) (string, error)
	Register(request users.UserEntity) error
}

type AuthDataInterface interface {
	GetUserByEmail(email string) (users.UserEntity, error)
	Register(request users.UserEntity) error
}
