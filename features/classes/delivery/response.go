package delivery

import "immersiveApp/features/users/delivery"

type ClassResponse struct {
	Id           uint                  `json:"id"`
	ClassName    string                `json:"class_name"`
	PicUserId    uint                  `json:"pic_user_id"`
	DateStart    string                `json:"date_start"`
	DateGraduate string                `json:"date_graduate"`
	User         delivery.UserResponse `json:"user,omitempty"`
}
