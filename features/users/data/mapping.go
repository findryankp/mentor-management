package data

import (
	"immersiveApp/features/teams"
	"immersiveApp/features/users"
	"reflect"
)

func UserEntityToUser(userEntity users.UserEntity) User {
	return User{
		TeamId:      userEntity.TeamId,
		FullName:    userEntity.FullName,
		Email:       userEntity.Email,
		Password:    userEntity.Password,
		PhoneNumber: userEntity.PhoneNumber,
		Address:     userEntity.Address,
		Role:        userEntity.Role,
		Status:      userEntity.Status,
	}
}

func UserToUserEntity(user User) users.UserEntity {

	result := users.UserEntity{
		Id:          user.ID,
		TeamId:      user.TeamId,
		FullName:    user.FullName,
		Email:       user.Email,
		Password:    user.Password,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
		Role:        user.Role,
		Status:      user.Status,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}

	if !reflect.ValueOf(user.Team).IsZero() {
		result.Team = teams.TeamEntity{
			Id:   user.Team.ID,
			Name: user.Team.Name,
		}
	}

	return result
}

func ListUserToUserEntity(user []User) []users.UserEntity {
	var userEntity []users.UserEntity
	for _, v := range user {
		userEntity = append(userEntity, UserToUserEntity(v))
	}
	return userEntity
}
