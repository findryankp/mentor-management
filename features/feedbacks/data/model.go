package data

import (
	// m "immersiveApp/features/mentees/data"
	s "immersiveApp/features/statuses/data"
	u "immersiveApp/features/users/data"

	"gorm.io/gorm"
)

type Feedback struct {
	gorm.Model
	Notes    string
	Proof    string
	UserId   uint
	User     *u.User `gorm:"foreignKey:UserId"`
	MenteeId uint
	// Mentee   *m.Mentee `gorm:"foreignKey:UserId"`
	// Mente m.Mentee `gorm:foreignKey:`
	StatusId uint
	Status   *s.Status `gorm:"foreignKey:UserId"`
	Approved bool
}

// type Result struct {
// 	m.Mentee
// 	Feedback
// }
