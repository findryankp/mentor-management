package statuses

import "time"

type StatusEntity struct {
	Id          uint
	Name        string `validate:"required"`
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type StatusServiceInterface interface {
	GetAll() ([]StatusEntity, error)
}

type StatusDataInterface interface {
	SelectAll() ([]StatusEntity, error)
}
