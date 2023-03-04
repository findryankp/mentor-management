package data

import "clean-arch/features/users"

func UserEntityToUser(userEntity users.UserEntity) User {
	return User{
		Name:     userEntity.Name,
		Email:    userEntity.Email,
		Password: userEntity.Password,
		Address:  userEntity.Address,
		Role:     userEntity.Role,
	}
}

func UserToUserEntity(user User) users.UserEntity {
	return users.UserEntity{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Address:   user.Address,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ListUserToUserEntity(user []User) []users.UserEntity {
	var userEntity []users.UserEntity
	for _, v := range user {
		userEntity = append(userEntity, UserToUserEntity(v))
	}
	return userEntity
}
