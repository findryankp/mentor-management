package delivery

type ClassRequest struct {
	ClassName    string `json:"class_name" form:"class_name"`
	PicUserId    uint   `json:"pic_user_id" form:"pic_user_id"`
	DateStart    string `json:"date_start" form:"date_start"`
	DateGraduate string `json:"date_graduate" form:"date_graduate"`
}
