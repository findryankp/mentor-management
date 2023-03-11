package data

import (
	"immersiveApp/features/teams"

	"gorm.io/gorm"
)

type query struct {
	db *gorm.DB
}

func New(db *gorm.DB) teams.TeamDataInterface {
	return &query{
		db: db,
	}
}

// SelectAll implements teams.TeamDataInterface
func (q *query) SelectAll() ([]teams.TeamEntity, error) {
	var teams []Team
	if err := q.db.Find(&teams); err.Error != nil {
		return nil, err.Error
	}
	return ListTeamToTeamEntity(teams), nil
}

// SelectById implements teams.TeamDataInterface
func (q *query) SelectById(id uint) (teams.TeamEntity, error) {
	var team Team
	if err := q.db.First(&team, id); err.Error != nil {
		return teams.TeamEntity{}, err.Error
	}
	return TeamToTeamEntity(team), nil
}

func (q *query) Store(teamEntity teams.TeamEntity) (uint, error) {
	team := TeamEntityToTeam(teamEntity)
	if err := q.db.Create(&team); err.Error != nil {
		return 0, err.Error
	}
	return team.ID, nil
}

func (q *query) Edit(teamEntity teams.TeamEntity, id uint) error {
	team := TeamEntityToTeam(teamEntity)
	if err := q.db.Where("id", id).Updates(&team); err.Error != nil {
		return err.Error
	}
	return nil
}

func (q *query) Destroy(id uint) error {
	var team Team
	if err := q.db.Delete(&team, id); err.Error != nil {
		return err.Error
	}
	return nil
}
