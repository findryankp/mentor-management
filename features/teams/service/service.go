package service

import (
	"immersiveApp/features/teams"

	"github.com/go-playground/validator/v10"
)

type teamService struct {
	Data     teams.TeamDataInterface
	validate *validator.Validate
}

func New(data teams.TeamDataInterface) teams.TeamServiceInterface {
	return &teamService{
		Data:     data,
		validate: validator.New(),
	}
}

func (s *teamService) GetAll() ([]teams.TeamEntity, error) {
	return s.Data.SelectAll()
}

func (s *teamService) GetById(id uint) (teams.TeamEntity, error) {
	return s.Data.SelectById(id)
}

func (s *teamService) Create(teamEntity teams.TeamEntity) (teams.TeamEntity, error) {
	errValidate := s.validate.Struct(teamEntity)
	if errValidate != nil {
		return teams.TeamEntity{}, errValidate
	}

	user_id, err := s.Data.Store(teamEntity)
	if err != nil {
		return teams.TeamEntity{}, err
	}
	return s.Data.SelectById(user_id)
}

func (s *teamService) Update(teamEntity teams.TeamEntity, id uint) (teams.TeamEntity, error) {
	//check exist data
	if checkDataExist, err := s.Data.SelectById(id); err != nil {
		return checkDataExist, err
	}

	// update
	err := s.Data.Edit(teamEntity, id)
	if err != nil {
		return teams.TeamEntity{}, err
	}
	return s.Data.SelectById(id)
}

func (s *teamService) Delete(id uint) error {
	return s.Data.Destroy(id)
}
