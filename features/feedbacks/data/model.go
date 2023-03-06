package data

import (
	m "clean-arch/features/mentees/data"
	s "clean-arch/features/statuses/data"
	u "clean-arch/features/users/data"

	"gorm.io/gorm"
)

type Feedback struct {
	gorm.Model
	Notes    string
	Proof    string
	UserId   int
	User     *u.User `gorm:"foreignKey:UserId"`
	MenteeId int
	Mentee   *m.Mentee `gorm:"foreignKey:UserId"`
	StatusId int
	Status   *s.Status `gorm:"foreignKey:UserId"`
	Approved bool
}
