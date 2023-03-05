package delivery

import (
	"clean-arch/features/teams/delivery"
	"clean-arch/features/users"
	"reflect"
)

type UserResponse struct {
	Id       uint                   `json:"id"`
	TeamId   uint                   `json:"team_id"`
	FullName string                 `json:"full_name"`
	Email    string                 `json:"email"`
	Role     string                 `json:"role"`
	Status   bool                   `json:"status"`
	Team     *delivery.TeamResponse `json:"team,omitempty"`
}

func UserEntityToUserResponse(userEntity users.UserEntity) UserResponse {
	result := UserResponse{
		Id:       userEntity.Id,
		TeamId:   userEntity.TeamId,
		FullName: userEntity.FullName,
		Email:    userEntity.Email,
		Role:     userEntity.Team.Name,
		Status:   userEntity.Status,
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
