package delivery

import (
	"immersiveApp/features/teams/delivery"
	"immersiveApp/features/users"
	"reflect"
)

type UserResponse struct {
	Id          uint                   `json:"id,omitempty"`
	TeamId      uint                   `json:"team_id,omitempty"`
	FullName    string                 `json:"full_name"`
	Email       string                 `json:"email,omitempty"`
	Role        string                 `json:"role,omitempty"`
	Status      bool                   `json:"status"`
	PhoneNumber string                 `json:"phone_number,omitempty"`
	Address     string                 `json:"address,omitempty"`
	Team        *delivery.TeamResponse `json:"team,omitempty"`
}

func UserEntityToUserResponse(userEntity users.UserEntity) UserResponse {
	result := UserResponse{
		Id:          userEntity.Id,
		TeamId:      userEntity.TeamId,
		FullName:    userEntity.FullName,
		Email:       userEntity.Email,
		Role:        userEntity.Role,
		Status:      userEntity.Status,
		PhoneNumber: userEntity.PhoneNumber,
		Address:     userEntity.Address,
	}

	if !reflect.ValueOf(userEntity.Team).IsZero() {
		result.Team = &delivery.TeamResponse{
			Id:   userEntity.Team.Id,
			Name: userEntity.Team.Name,
		}
	}

	return result
}

func ListUserEntityToUserResponse(dataCore []users.UserEntity) []UserResponse {
	var dataResponses []UserResponse
	for _, v := range dataCore {
		dataResponses = append(dataResponses, UserEntityToUserResponse(v))
	}
	return dataResponses
}
