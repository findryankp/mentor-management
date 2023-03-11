package data

import (
	"immersiveApp/features/feedbacks"

	"gorm.io/gorm"
)

type query struct {
	db *gorm.DB
}

func New(db *gorm.DB) feedbacks.FeedbackDataInterface {
	return &query{
		db: db,
	}
}

func (q *query) SelectAll() ([]feedbacks.FeedbackEntity, error) {
	var feedbacks []Feedback
	if err := q.db.Find(&feedbacks); err.Error != nil {
		return nil, err.Error
	}
	return ListFeedbackToFeedbackEntity(feedbacks), nil
}

func (q *query) SelectById(id uint) (feedbacks.FeedbackEntity, error) {
	var feedback Feedback
	if err := q.db.First(&feedback, id); err.Error != nil {
		return feedbacks.FeedbackEntity{}, err.Error
	}
	return FeedbackToFeedbackEntity(feedback), nil
}

func (q *query) Insert(feedbackEntity feedbacks.FeedbackEntity) (uint, error) {
	feedback := FeedbackEntityToFeedback(feedbackEntity)
	if err := q.db.Create(&feedback); err.Error != nil {
		return 0, err.Error
	}
	return feedback.ID, nil
}

func (q *query) Edit(feedbackEntity feedbacks.FeedbackEntity, id uint) error {
	feedback := FeedbackEntityToFeedback(feedbackEntity)
	if err := q.db.Where("id", id).Updates(&feedback); err.Error != nil {
		return err.Error
	}
	return nil
}

func (q *query) Remove(id uint) error {
	var feedback Feedback
	if err := q.db.Delete(&feedback, id); err.Error != nil {
		return err.Error
	}
	return nil
}
