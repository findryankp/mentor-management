package mentees

import (
	"immersiveApp/features/classes"
	// "immersiveApp/features/feedbacks"

	// "immersiveApp/features/feedbacks"

	"immersiveApp/features/statuses"

	"time"
)

type MenteeEntity struct {
	Id              uint
	ClassId         uint `validate:"required"`
	Class           classes.ClassEntity
	FullName        string `validate:"required"`
	NickName        string `validate:"required"`
	Email           string `validate:"required"`
	Phone           string `validate:"required"`
	CurrentAddress  string `validate:"required"`
	HomeAddress     string `validate:"required"`
	Telegram        string `validate:"required"`
	StatusId        uint   `validate:"required"`
	Status          statuses.StatusEntity
	Gender          string `validate:"required"`
	EducationType   string `validate:"required"`
	Major           string `validate:"required"`
	Graduate        string `validate:"required"`
	Institution     string `validate:"required"`
	EmergencyName   string `validate:"required"`
	EmergencyPhone  string `validate:"required"`
	EmergencyStatus string `validate:"required"`
	// Feedbacks []string
	// Feedback        feedbacks.FeedbackEntity
	CreatedAt       time.Time
	UpdatedAt       time.Time
}


type MenteeServiceInterface interface {
	GetAll() ([]MenteeEntity, error)
	GetById(id uint) (MenteeEntity, error)
	GetFeedbackById(id uint) (any, error)
	Create(menteeEntity MenteeEntity) (MenteeEntity, error)
	Update(menteeEntity MenteeEntity, id uint) (MenteeEntity, error)
	Delete(id uint) error
}

type MenteeDataInterface interface {
	SelectAll() ([]MenteeEntity, error)
	SelectById(id uint) (MenteeEntity, error)
	SelectFeedbackById(id uint) (any, error)
	Store(menteeEntity MenteeEntity) (uint, error)
	Edit(menteeEntity MenteeEntity, id uint) error
	Destroy(id uint) error
}
