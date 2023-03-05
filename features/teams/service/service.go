package service

import (
	"clean-arch/features/teams"
)

type teamService struct {
	Data teams.TeamDataInterface
}

func New(data teams.TeamDataInterface) teams.TeamServiceInterface {
	return &teamService{
		Data: data,
	}
}

func (t *teamService) GetAll() ([]teams.TeamEntity, error) {
	return t.Data.SelectAll()
}

func (t *teamService) GetById(id uint) (teams.TeamEntity, error) {
	return t.Data.SelectById(id)
}

// Create implements teams.TeamServiceInterface
func (t *teamService) Create(teamEntity teams.TeamEntity) (teams.TeamEntity, error) {
	return t.Data.Store(teamEntity)
}

func (t *teamService) Update(teamEntity teams.TeamEntity, id uint) (teams.TeamEntity, error) {
	//check exist data
	if checkDataExist, err := t.Data.SelectById(id); err != nil {
		return checkDataExist, err
	}

	// update
	updateData, err := t.Data.Edit(teamEntity, id)
	updateData.Id = id
	return updateData, err
}

func (t *teamService) Delete(id uint) error {
	return t.Data.Destroy(id)
}
