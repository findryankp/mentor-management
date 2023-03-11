package data

import (
	"immersiveApp/features/classes"

	"gorm.io/gorm"
)

type query struct {
	db *gorm.DB
}

func New(db *gorm.DB) classes.ClassDataInterface {
	return &query{
		db: db,
	}
}

func (q *query) SelectAll() ([]classes.ClassEntity, error) {
	var classes []Class
	if err := q.db.Preload("User").Find(&classes); err.Error != nil {
		return nil, err.Error
	}
	return ListClassToClassEntity(classes), nil
}

func (q *query) SelectById(id uint) (classes.ClassEntity, error) {
	var class Class
	if err := q.db.Preload("User").First(&class, id); err.Error != nil {
		return classes.ClassEntity{}, err.Error
	}
	return ClassToClassEntity(class), nil
}

func (q *query) Store(classEntity classes.ClassEntity) (uint, error) {
	class := ClassEntityToClass(classEntity)
	if err := q.db.Create(&class); err.Error != nil {
		return 0, err.Error
	}
	return class.ID, nil
}

func (q *query) Edit(classEntity classes.ClassEntity, id uint) error {
	class := ClassEntityToClass(classEntity)
	if err := q.db.Where("id", id).Updates(&class); err.Error != nil {
		return err.Error
	}
	return nil
}

func (q *query) Destroy(id uint) error {
	var class Class
	if err := q.db.Delete(&class, id); err.Error != nil {
		return err.Error
	}
	return nil
}
