package service

import (
	"errors"
	"immersiveApp/features/mentees"
	"immersiveApp/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	repo := mocks.NewMenteeDataInterface(t)
	input := mentees.MenteeEntity{
		Id:              uint(1),
		ClassId:         1,
		FullName:        "Jomabala",
		NickName:        "Balvir",
		Email:           "bal@hotmail.com",
		Phone:           "089",
		CurrentAddress:  "jl santai",
		HomeAddress:     "jl rusak",
		Telegram:        "@siapakamu",
		StatusId:        1,
		Gender:          "P",
		EducationType:   "non-informatics",
		Major:           "Chemistry",
		Graduate:        "1030",
		Institution:     "UniGa",
		EmergencyName:   "Blla",
		EmergencyPhone:  "089",
		EmergencyStatus: "Ibu",
	}
	srv := New(repo)

	mockinput := mentees.MenteeEntity{
		Id:              uint(2),
		ClassId:         1,
		FullName:        "",
		NickName:        "",
		Email:           "",
		Phone:           "",
		CurrentAddress:  "",
		HomeAddress:     "",
		Telegram:        "",
		StatusId:        1,
		Gender:          "",
		EducationType:   "",
		Major:           "",
		Graduate:        "",
		Institution:     "",
		EmergencyName:   "",
		EmergencyPhone:  "",
		EmergencyStatus: "",
	}
	t.Run("succes", func(t *testing.T) {
		repo.On("Store", input).Return(uint(1), nil)
		repo.On("SelectById", uint(1)).Return(input, nil)

		created, err := srv.Create(input)

		assert.NoError(t, err)
		assert.Equal(t, input, created)

		repo.AssertExpectations(t)
	})

	t.Run("eror", func(t *testing.T) {
		_, err := srv.Create(mockinput)
		assert.NotEmpty(t, err)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestGetAll(t *testing.T) {
	repo := mocks.NewMenteeDataInterface(t)
	input := mentees.MenteeEntity{
		Id:              uint(1),
		ClassId:         1,
		FullName:        "Jomabala",
		NickName:        "Balvir",
		Email:           "bal@hotmail.com",
		Phone:           "089",
		CurrentAddress:  "jl santai",
		HomeAddress:     "jl rusak",
		Telegram:        "@siapakamu",
		StatusId:        1,
		Gender:          "P",
		EducationType:   "non-informatics",
		Major:           "Chemistry",
		Graduate:        "1030",
		Institution:     "UniGa",
		EmergencyName:   "Blla",
		EmergencyPhone:  "089",
		EmergencyStatus: "Ibu",
	}
	srv := New(repo)

	mockList := make([]mentees.MenteeEntity, 0)
	mockList = append(mockList, input)

	repo.On("SelectAll").Return(mockList, nil)

	result, err := srv.GetAll()

	repo.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Len(t, result, len(mockList))
	assert.Equal(t, mockList, result)
}

func TestGetById(t *testing.T) {
	repo := mocks.NewMenteeDataInterface(t)
	input := mentees.MenteeEntity{
		Id:              uint(1),
		ClassId:         1,
		FullName:        "meong milanium",
		NickName:        "meong",
		Email:           "miyong@mail",
		Phone:           "089765",
		CurrentAddress:  "jl. sesama",
		HomeAddress:     "malang",
		Telegram:        "youma",
		StatusId:        1,
		Gender:          "L",
		EducationType:   "non-informatics",
		Major:           "DKV",
		Graduate:        "1040",
		Institution:     "UniGa",
		EmergencyName:   "Jola",
		EmergencyPhone:  "089890",
		EmergencyStatus: "Kakek",
	}
	repo.On("SelectById", input.Id).Return(input, nil)

	srv := New(repo)

	result, err := srv.GetById(input.Id)
	repo.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, input, result)
}

func TestGetFeedbackById(t *testing.T) {
	repo := mocks.NewMenteeDataInterface(t)
	input := mentees.MenteeEntity{
		Id:              uint(1),
		ClassId:         1,
		FullName:        "meong milanium",
		NickName:        "meong",
		Email:           "miyong@mail",
		Phone:           "089765",
		CurrentAddress:  "jl. sesama",
		HomeAddress:     "malang",
		Telegram:        "youma",
		StatusId:        1,
		Gender:          "L",
		EducationType:   "non-informatics",
		Major:           "DKV",
		Graduate:        "1040",
		Institution:     "UniGa",
		EmergencyName:   "Jola",
		EmergencyPhone:  "089890",
		EmergencyStatus: "Kakek",
	}
	repo.On("SelectFeedbackById", input.Id).Return(input, nil)

	srv := New(repo)

	result, err := srv.GetFeedbackById(input.Id)
	repo.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, input, result)
}

func TestUpdate(t *testing.T) {
	repo := mocks.NewMenteeDataInterface(t)

	input := mentees.MenteeEntity{
		Id:              uint(1),
		ClassId:         1,
		FullName:        "meong milanium",
		NickName:        "meong",
		Email:           "miyong@mail",
		Phone:           "089765",
		CurrentAddress:  "jl. sesama",
		HomeAddress:     "malang",
		Telegram:        "youma",
		StatusId:        1,
		Gender:          "L",
		EducationType:   "non-informatics",
		Major:           "DKV",
		Graduate:        "1040",
		Institution:     "UniGa",
		EmergencyName:   "Jola",
		EmergencyPhone:  "089890",
		EmergencyStatus: "Kakek",
	}
	updatedData := mentees.MenteeEntity{
		Id:              uint(1),
		ClassId:         1,
		FullName:        "meong milanium",
		NickName:        "meong",
		Email:           "miyong@mail",
		Phone:           "089765",
		CurrentAddress:  "jl. sesama",
		HomeAddress:     "malang",
		Telegram:        "youma",
		StatusId:        1,
		Gender:          "L",
		EducationType:   "non-informatics",
		Major:           "DKV",
		Graduate:        "1040",
		Institution:     "UniGa",
		EmergencyName:   "Jola",
		EmergencyPhone:  "089890",
		EmergencyStatus: "Kakek",
	}
	t.Run("sukses update data", func(t *testing.T) {
		repo.On("Update", uint(1), input).Return(updatedData, nil).Once()

		s := &MenteeService{Data: repo}

		updated, err := s.Update(updatedData, input.Id)

		assert.Nil(t, err)
		assert.Equal(t, updatedData, updated)
	})
	t.Run("data tidak ditemukan", func(t *testing.T) {
		repo.On("Update", uint(3), input).Return(updatedData, errors.New("data not found")).Once()

		service := New(repo)
		res, err := service.Update(updatedData, input.Id)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak ditemukan")
		assert.Equal(t, uint(0), res.Id)
		repo.AssertExpectations(t)
	})
	t.Run("eror", func(t *testing.T) {
		service := New(repo)
		_, err := service.Update(updatedData, input.Id)

		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "eror")
	})
}

func TestDelete(t *testing.T) {
	repo := mocks.NewMenteeDataInterface(t)
	input := mentees.MenteeEntity{
		Id:              uint(1),
		ClassId:         1,
		FullName:        "meong milanium",
		NickName:        "meong",
		Email:           "miyong@mail",
		Phone:           "089765",
		CurrentAddress:  "jl. sesama",
		HomeAddress:     "malang",
		Telegram:        "youma",
		StatusId:        1,
		Gender:          "L",
		EducationType:   "non-informatics",
		Major:           "DKV",
		Graduate:        "1040",
		Institution:     "UniGa",
		EmergencyName:   "Jola",
		EmergencyPhone:  "089890",
		EmergencyStatus: "Kakek",
	}
	t.Run("Sukses Delete", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(input, nil).Once()
		repo.On("SelectById", uint(1)).Return(input, nil)
		srv := New(repo)
		err := srv.Delete(1)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Gagal Delete", func(t *testing.T) {
		repo.On("SelectById", uint(1)).Return(input, nil)
		srv := New(repo)
		err := srv.Delete(uint(0))
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}
