package service

import (
	"immersiveApp/features/mentees"

	"github.com/go-playground/validator/v10"
)

type MenteeService struct {
	Data     mentees.MenteeDataInterface
	validate *validator.Validate
}

func New(data mentees.MenteeDataInterface) mentees.MenteeServiceInterface {
	return &MenteeService{
		Data:     data,
		validate: validator.New(),
	}
}

func (s *MenteeService) GetAll() ([]mentees.MenteeEntity, error) {
	return s.Data.SelectAll()
}

func (s *MenteeService) GetById(id uint) (mentees.MenteeEntity, error) {
	return s.Data.SelectById(id)
}

func (s *MenteeService) GetFeedbackById(id uint) (any, error) {
	return s.Data.SelectFeedbackById(id)
}

func (s *MenteeService) Create(menteeEntity mentees.MenteeEntity) (mentees.MenteeEntity, error) {
	s.validate = validator.New()
	errValidate := s.validate.StructExcept(menteeEntity, "Class", "Status")
	if errValidate != nil {
		return mentees.MenteeEntity{}, errValidate
	}

	user_id, err := s.Data.Store(menteeEntity)
	if err != nil {
		return mentees.MenteeEntity{}, err
	}

	return s.Data.SelectById(user_id)
}

func (s *MenteeService) Update(menteeEntity mentees.MenteeEntity, id uint) (mentees.MenteeEntity, error) {
	if checkDataExist, err := s.Data.SelectById(id); err != nil {
		return checkDataExist, err
	}

	err := s.Data.Edit(menteeEntity, id)
	if err != nil {
		return mentees.MenteeEntity{}, err
	}
	return s.Data.SelectById(id)
}

func (s *MenteeService) Delete(id uint) error {
	if _, err := s.Data.SelectById(id); err != nil {
		return err
	}

	return s.Data.Destroy(id)
}
