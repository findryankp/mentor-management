package delivery

type FeedbackResponse struct {
	Id       uint `json:"id"`
	Notes    string `json:"notes"`
	Proof    string `json:"proof"`
	UserId   uint `json:"user_id"`
	MenteeId uint `json:"mentee_id"`
	StatusId uint `json:"status_id"`
	Approved bool `json:"approved"`
}