package delivery

type FeedbackRequest struct {
	Notes    string `json:"notes" form:"notes"`
	Proof    string `json:"proof" form:"proof"`
	UserId   uint `json:"user_id" form:"user_id"`
	MenteeId uint `json:"mentee_id" form:"mentee_id"`
	StatusId uint `json:"status_id" form:"status_id"`
	Approved bool `json:"approved" form:"approved"`
}