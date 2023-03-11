package service

import (
	"errors"
	"immersiveApp/features/classes"
	"immersiveApp/utils/helpers"

	"github.com/go-playground/validator/v10"
)

type ClassService struct {
	Data     classes.ClassDataInterface
	validate *validator.Validate
}

func New(data classes.ClassDataInterface) classes.ClassServiceInterface {
	return &ClassService{
		Data:     data,
		validate: validator.New(),
	}
}

func (s *ClassService) GetAll() ([]classes.ClassEntity, error) {
	return s.Data.SelectAll()
}

func (s *ClassService) GetById(id uint) (classes.ClassEntity, error) {
	return s.Data.SelectById(id)
}

func (s *ClassService) Create(classEntity classes.ClassEntity) (classes.ClassEntity, error) {
	s.validate = validator.New()
	errValidate := s.validate.StructExcept(classEntity, "User")
	if errValidate != nil {
		return classes.ClassEntity{}, errValidate
	}

	dateStart, checkDateStart := helpers.IsDate(classEntity.DateStart)
	if !checkDateStart {
		return classes.ClassEntity{}, errors.New("not valid date start format date, ex : 2006-02-25")
	}

	dateGraduate, checkDateGraduate := helpers.IsDate(classEntity.DateGraduate)
	if !checkDateGraduate {
		return classes.ClassEntity{}, errors.New("not valid date graduate format date, ex : 2006-02-25")
	}

	if helpers.FormatDate(dateGraduate).Before(helpers.FormatDate(dateStart)) {
		return classes.ClassEntity{}, errors.New("error range of date, date graduate must be after date start")
	}

	user_id, err := s.Data.Store(classEntity)
	if err != nil {
		return classes.ClassEntity{}, err
	}

	return s.Data.SelectById(user_id)
}

func (s *ClassService) Update(classEntity classes.ClassEntity, id uint) (classes.ClassEntity, error) {
	if checkDataExist, err := s.Data.SelectById(id); err != nil {
		return checkDataExist, err
	}

	err := s.Data.Edit(classEntity, id)
	if err != nil {
		return classes.ClassEntity{}, err
	}
	return s.Data.SelectById(id)
}

func (s *ClassService) Delete(id uint) error {
	if _, err := s.Data.SelectById(id); err != nil {
		return err
	}

	return s.Data.Destroy(id)
}
