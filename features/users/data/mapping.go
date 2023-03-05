package data

import (
	"clean-arch/features/teams"
	"clean-arch/features/users"
)

func UserEntityToUser(userEntity users.UserEntity) User {
	return User{
		TeamId:   userEntity.TeamId,
		FullName: userEntity.FullName,
		Email:    userEntity.Email,
		Password: userEntity.Password,
		Role:     userEntity.Role,
		Status:   userEntity.Status,
	}
}

func UserToUserEntity(user User) users.UserEntity {
	return users.UserEntity{
		Id:        user.ID,
		TeamId:    user.TeamId,
		FullName:  user.FullName,
		Email:     user.Email,
		Password:  user.Password,
		Role:      user.Role,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Team: teams.TeamEntity{
			Id:   user.Team.ID,
			Name: user.Team.Name,
		},
	}
}

func ListUserToUserEntity(user []User) []users.UserEntity {
	var userEntity []users.UserEntity
	for _, v := range user {
		userEntity = append(userEntity, UserToUserEntity(v))
	}
	return userEntity
}
