package data

import (
	"immersiveApp/features/statuses"
	status "immersiveApp/features/statuses"

	"gorm.io/gorm"
)

type query struct {
	db *gorm.DB
}

func New(db *gorm.DB) status.StatusDataInterface {
	return &query{
		db: db,
	}
}

// SelectAll implements status.StatusDataInterface
func (q *query) SelectAll() ([]statuses.StatusEntity, error) {
	var status []Status
	if err := q.db.Find(&status); err.Error != nil {
		return nil, err.Error
	}
	return ListStatusToStatusEntity(status), nil
}
