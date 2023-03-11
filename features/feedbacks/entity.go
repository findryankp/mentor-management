package feedbacks

import (
	"immersiveApp/features/statuses"
	"immersiveApp/features/users"
	"time"
)

type FeedbackEntity struct {
	Id        uint
	Notes     string `validate:"required"`
	Proof     string `validate:"required"`
	UserId    uint
	User      users.UserEntity
	MenteeId  uint
	StatusId  uint
	Status    statuses.StatusEntity
	Approved  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type FeedbackServiceInterface interface {
	GetAll() ([]FeedbackEntity, error)
	GetById(id uint) (FeedbackEntity, error)
	Create(feedbackEntity FeedbackEntity) (FeedbackEntity, error)
	Update(feedbackEntity FeedbackEntity, id uint) (FeedbackEntity, error)
	Delete(id uint) error
}

type FeedbackDataInterface interface {
	SelectAll() ([]FeedbackEntity, error)
	SelectById(id uint) (FeedbackEntity, error)
	Insert(feedbackEntity FeedbackEntity) (uint, error)
	Edit(feedbackEntity FeedbackEntity, id uint) error
	Remove(id uint) error
}
