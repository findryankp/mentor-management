package teams

import (
	"time"
)

type TeamEntity struct {
	Id        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TeamServiceInterface interface {
	GetAll() ([]TeamEntity, error)
	GetById(id uint) (TeamEntity, error)
	Create(teamEntity TeamEntity) (TeamEntity, error)
	Update(teamEntity TeamEntity, id uint) (TeamEntity, error)
	Delete(id uint) error
}

type TeamDataInterface interface {
	SelectAll() ([]TeamEntity, error)
	SelectById(id uint) (TeamEntity, error)
	Store(teamEntity TeamEntity) (TeamEntity, error)
	Edit(teamEntity TeamEntity, id uint) (TeamEntity, error)
	Destroy(id uint) error
}
