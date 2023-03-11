package service

import (
	"immersiveApp/features/feedbacks"

	"github.com/go-playground/validator/v10"
)

type feedbackService struct {
	Data     feedbacks.FeedbackDataInterface
	validate *validator.Validate
}

func New(data feedbacks.FeedbackDataInterface) feedbacks.FeedbackServiceInterface {
	return &feedbackService{
		Data:     data,
		validate: validator.New(),
	}
}

func (s *feedbackService) GetAll() ([]feedbacks.FeedbackEntity, error) {
	return s.Data.SelectAll()
}

func (s *feedbackService) GetById(id uint) (feedbacks.FeedbackEntity, error) {
	return s.Data.SelectById(id)
}

func (s *feedbackService) Create(feedbackEntity feedbacks.FeedbackEntity) (feedbacks.FeedbackEntity, error) {
	errValidate := s.validate.StructExcept(feedbackEntity, "User", "Status")
	if errValidate != nil {
		return feedbacks.FeedbackEntity{}, errValidate
	}
	feedbackEntity.Approved = true
	user_id, err := s.Data.Insert(feedbackEntity)
	if err != nil {
		return feedbacks.FeedbackEntity{}, err
	}
	return s.Data.SelectById(user_id)
}

func (s *feedbackService) Update(feedbackEntity feedbacks.FeedbackEntity, id uint) (feedbacks.FeedbackEntity, error) {
	//check exist data
	if checkDataExist, err := s.Data.SelectById(id); err != nil {
		return checkDataExist, err
	}

	//update
	err := s.Data.Edit(feedbackEntity, id)
	if err != nil {
		return feedbacks.FeedbackEntity{}, err
	}
	return s.Data.SelectById(id)
}

func (s *feedbackService) Delete(id uint) error {
	if _, err := s.Data.SelectById(id); err != nil {
		return err
	}
	return s.Data.Remove(id)
}
