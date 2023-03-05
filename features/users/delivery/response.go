package delivery

import "clean-arch/features/users"

type UserResponse struct {
	Id       uint   `json:"id"`
	TeamId   uint   `json:"team_id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Status   bool   `json:"status"`
}

func UserEntityToUserResponse(userEntity users.UserEntity) UserResponse {
	return UserResponse{
		Id:       userEntity.Id,
		TeamId:   userEntity.TeamId,
		FullName: userEntity.FullName,
		Email:    userEntity.Email,
		Role:     userEntity.Role,
		Status:   userEntity.Status,
	}
}

func ListUserEntityToUserResponse(dataCore []users.UserEntity) []UserResponse {
	var dataResponses []UserResponse
	for _, v := range dataCore {
		dataResponses = append(dataResponses, UserEntityToUserResponse(v))
	}
	return dataResponses
}
