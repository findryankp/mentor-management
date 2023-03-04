package data

import (
	"clean-arch/features/users"

	"gorm.io/gorm"
)

type query struct {
	db *gorm.DB
}

func New(db *gorm.DB) users.UserDataInterface {
	return &query{
		db: db,
	}
}

func (q *query) SelectAll() ([]users.UserEntity, error) {
	var users []User
	if err := q.db.Find(&users); err.Error != nil {
		return nil, err.Error
	}
	return ListUserToUserEntity(users), nil
}

func (q *query) SelectById(id uint) (users.UserEntity, error) {
	var user User
	if err := q.db.First(&user, id); err.Error != nil {
		return users.UserEntity{}, err.Error
	}
	return UserToUserEntity(user), nil
}

func (q *query) Store(request users.UserEntity) error {
	user := UserEntityToUser(request)
	if err := q.db.Create(&user); err.Error != nil {
		return err.Error
	}
	return nil
}

func (q *query) Edit(request users.UserEntity, id uint) error {
	user := UserEntityToUser(request)
	if err := q.db.Where("id", id).Updates(&user); err.Error != nil {
		return err.Error
	}
	return nil
}

func (q *query) Destroy(id uint) error {
	var user User
	if err := q.db.Delete(&user, id); err.Error != nil {
		return err.Error
	}
	return nil
}
