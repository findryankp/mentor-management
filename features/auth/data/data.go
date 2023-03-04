package data

import (
	"clean-arch/features/auth"
	"clean-arch/features/users"
	"clean-arch/features/users/data"

	"gorm.io/gorm"
)

type query struct {
	db *gorm.DB
}

// Login implements authuser.AuthDataInterface
func (q *query) GetUserByEmail(email string) (users.UserEntity, error) {
	var user data.User
	if err := q.db.Where("email", email).First(&user); err.Error != nil {
		return users.UserEntity{}, err.Error
	}

	return data.UserToUserEntity(user), nil
}

func (q *query) Register(request users.UserEntity) error {
	user := data.UserEntityToUser(request)
	if err := q.db.Create(&user); err.Error != nil {
		return err.Error
	}
	return nil
}

func New(db *gorm.DB) auth.AuthDataInterface {
	return &query{
		db: db,
	}
}
