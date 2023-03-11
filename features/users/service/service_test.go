package service

import (
	"errors"
	"immersiveApp/features/users"
	"immersiveApp/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	repo := mocks.NewUserDataInterface(t)

	input2 := users.UserEntity{
		Id:          uint(1),
		TeamId:      1,
		FullName:    "rischi",
		Email:       "rischi@mail",
		Password:    "12345",
		PhoneNumber: "08123",
		Address:     "pekanbaru",
		Role:        "active",
		Status:      false,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
	srv := New(repo)

	mockinput := users.UserEntity{
		Id:          uint(2),
		TeamId:      1,
		FullName:    "a",
		Email:       "a@mail",
		Password:    "a",
		PhoneNumber: "123",
		Address:     "a",
		Role:        "arole",
		Status:      false,
	}

	t.Run("success", func(t *testing.T) {
		repo.On("Store", input2).Return(uint(1), nil)
		repo.On("SelectById", uint(1)).Return(input2, nil)
		created, err := srv.Create(input2)

		assert.NoError(t, err)
		assert.Equal(t, input2, created)
		repo.AssertExpectations(t)
	})

	t.Run("data not found", func(t *testing.T) {
		mockinput := users.UserEntity{
			Id:          uint(2),
			TeamId:      1,
			FullName:    "",
			Email:       "",
			Password:    "",
			PhoneNumber: "",
			Address:     "",
			Role:        "",
			Status:      false,
		}
		email := "riau@gmail.com"
		repo.On("Login", email).Return(users.UserEntity{}, errors.New("data not found"))

		srv := New(repo)
		created, err := srv.Create(mockinput)

		assert.Empty(t, created)
		assert.EqualError(t, err, "data not found")
		repo.AssertExpectations(t)
	})
	t.Run("eror", func(t *testing.T) {
		_, err := srv.Create(mockinput)
		assert.NotEmpty(t, err)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("role option only : admin and user", func(t *testing.T) {

		_, err := srv.Create(input2)

		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "role option only : admin and user")
	})
}

func TestGetAll(t *testing.T) {
	repo := mocks.NewUserDataInterface(t)
	input := users.UserEntity{
		Id:          uint(1),
		TeamId:      1,
		FullName:    "rischi",
		Email:       "rischi@mail",
		Password:    "12345",
		PhoneNumber: "08123",
		Address:     "pekanbaru",
		Role:        "active",
		Status:      false,
	}
	srv := New(repo)

	mockList := make([]users.UserEntity, 0)
	mockList = append(mockList, input)

	repo.On("SelectAll").Return(mockList, nil)

	result, err := srv.GetAll()

	repo.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Len(t, result, len(mockList))
	assert.Equal(t, mockList, result)

}
func TestGetById(t *testing.T) {
	repo := mocks.NewUserDataInterface(t)
	input := users.UserEntity{
		Id:          uint(1),
		TeamId:      1,
		FullName:    "rischi",
		Email:       "rischi@mail",
		Password:    "12345",
		PhoneNumber: "08123",
		Address:     "pekanbaru",
		Role:        "active",
		Status:      false,
	}

	repo.On("SelectById", input.Id).Return(input, nil)

	srv := New(repo)

	result, err := srv.GetById(input.Id)
	repo.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, input, result)
}

func TestDelete(t *testing.T) {
	repo := mocks.NewUserDataInterface(t)
	input := users.UserEntity{
		Id:          uint(1),
		FullName:    "rischi",
		Email:       "rischi@mail",
		Password:    "12345",
		PhoneNumber: "08123",
		Address:     "pekanbaru",
		Role:        "active",
		Status:      false,
	}
	t.Run("Sukses Delete", func(t *testing.T) {
		repo.On("Delete", uint(1)).Return(users.UserEntity{}, nil).Once()
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

func TestUpdate(t *testing.T) {
	repo := mocks.NewUserDataInterface(t)

	input := users.UserEntity{
		Id:          uint(1),
		FullName:    "rischi",
		Email:       "rischi@mail",
		Password:    "12345",
		PhoneNumber: "08123",
		Address:     "pekanbaru",
		Role:        "user",
		Status:      true,
	}
	updatedData := users.UserEntity{
		Id:          1,
		FullName:    "rischi update",
		Email:       "rischi@mail",
		Password:    "12345",
		PhoneNumber: "08123",
		Address:     "pekanbaru",
		Role:        "user",
		Status:      true,
	}
	t.Run("sukses update data", func(t *testing.T) {
		repo.On("Update", uint(1), input).Return(updatedData, nil).Once()

		s := &userService{Data: repo}

		updated, err := s.Update(updatedData, input.Id)

		assert.Nil(t, err)
		assert.Equal(t, updatedData, updated)
	})
	t.Run("data tidak ditemukan", func(t *testing.T) {
		repo.On("Update", uint(5), input).Return(updatedData, errors.New("data not found")).Once()

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
