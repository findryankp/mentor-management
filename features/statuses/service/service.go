package service

import (
	status "immersiveApp/features/statuses"

	"github.com/go-playground/validator/v10"
)

type statusService struct {
	Data     status.StatusDataInterface
	validate *validator.Validate
}

func New(data status.StatusDataInterface) status.StatusServiceInterface {
	return &statusService{
		Data:     data,
		validate: validator.New(),
	}
}

func (s *statusService) GetAll() ([]status.StatusEntity, error) {
	return s.Data.SelectAll()
}
