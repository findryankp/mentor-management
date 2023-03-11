package delivery

type MenteeRequest struct {
	ClassId        uint   `json:"class_id" form:"class_id"`
	FullName       string `json:"full_name" form:"full_name"`
	NickName       string `json:"nick_name" form:"nick_name"`
	Email          string `json:"email" form:"email"`
	Phone          string `json:"phone" form:"phone"`
	CurrentAddress string `json:"current_address" form:"current_address"`
	HomeAddress    string `json:"home_address" form:"home_address"`
	Telegram       string `json:"telegram" form:"telegram"`
	StatusId       uint   `json:"status_id" form:"status_id"`
	// Status          *delivery.StatusResponse
	Gender          string `json:"gender" form:"gender"`
	EducationType   string `json:"education_type" form:"education_type"`
	Major           string `json:"major" form:"major"`
	Graduate        string `json:"graduate" form:"graduate"`
	Institution     string `json:"institution" form:"institution"`
	EmergencyName   string `json:"emergency_name" form:"emergency_name"`
	EmergencyPhone  string `json:"emergency_phone" form:"emergency_phone"`
	EmergencyStatus string `json:"emergency_status" form:"emergency_status"`
}
