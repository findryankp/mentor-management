package delivery

import "clean-arch/features/users"

type UserRequest struct {
	TeamId   uint   `json:"team_id" form:"team_id"`
	FullName string `json:"full_name" form:"full_name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Role     string `json:"role" form:"role"`
	Status   bool   `json:"status" form:"status"`
}

func UserRequestToUserEntity(userRequest UserRequest) users.UserEntity {
	return users.UserEntity{
		TeamId:   userRequest.TeamId,
		FullName: userRequest.FullName,
		Email:    userRequest.Email,
		Password: userRequest.Password,
		Role:     userRequest.Role,
		Status:   userRequest.Status,
	}
}
