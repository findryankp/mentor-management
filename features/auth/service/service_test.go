package service

import (
	"errors"
	"immersiveApp/features/users"
	"immersiveApp/mocks"
	"immersiveApp/utils/helpers"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLogin(t *testing.T) {
	data := mocks.NewAuthDataInterface(t)
	srv := New(data)
	hash, _ := helpers.HashPassword("12345")
	input := users.UserEntity{
		Id:          uint(1),
		TeamId:      1,
		FullName:    "rischi",
		Email:       "rischi@mail",
		Password:    hash,
		PhoneNumber: "0812345",
		Address:     "pekanbaru",
		Role:        "active",
		Status:      false,
	}
	input2 := users.UserEntity{
		Id:          uint(0),
		TeamId:      0,
		FullName:    "",
		Email:       "",
		Password:    "12345",
		PhoneNumber: "",
		Address:     "",
		Role:        "",
		Status:      false,
	}
	t.Run("succes", func(t *testing.T) {
		data.On("Store", input).Return(uint(1), nil)
		data.On("SelectById", uint(1)).Return(input, nil)

		_, err := srv.Login(input.Email, input.Password)

		assert.NoError(t, err)
		assert.Equal(t, input, err)

		data.AssertExpectations(t)
	})
	t.Run("eror", func(t *testing.T) {
		_, err := srv.Login(input2.Email, input2.Password)
		assert.NotEmpty(t, err)
		assert.NotNil(t, err)
		data.AssertExpectations(t)
	})

	t.Run("email and password must be filled", func(t *testing.T) {
		token, err := srv.Login("", "")
		assert.EqualError(t, err, "email and password must be fill")
		assert.Empty(t, token)
	})

}

func TestAuthService_ChangePassword(t *testing.T) {
	data := mocks.NewAuthDataInterface(t)
	srv := New(data)
	password := "password123"
	hashedPassword, _ := helpers.HashPassword(password)

	data.On("GetUserByEmailOrId", ".", uint(1)).Return(&users.UserEntity{
		Id:       1,
		Email:    "john@example.com",
		Password: hashedPassword,
	}, nil)

	data.On("EditPassword", uint(1), mock.AnythingOfType("string")).Return(nil)

	err := srv.ChangePassword(1, password, password, password)
	assert.NoError(t, err)

	data.AssertExpectations(t)

	t.Run("validation eror", func(t *testing.T) {
		err := srv.ChangePassword(1, "", "", "")
		assert.EqualError(t, err, "old password,new password, and confirm password cannot be empty")

		err = srv.ChangePassword(1, "oldPassword", "newPassword", "confirmPassword")
		assert.EqualError(t, err, "new password and confirm password must be similarity")

	})

	t.Run("Fail", func(t *testing.T) {
		data.On("EditPassword", uint(1), mock.AnythingOfType("string")).Return(errors.New("error"))

		err := srv.ChangePassword(1, password, password, password)
		assert.EqualError(t, err, "failed to change password")

		data.AssertExpectations(t)
	})
	t.Run("eror", func(t *testing.T) {
		err := srv.ChangePassword(1, "", "", "")
		assert.NotEmpty(t, err)
		assert.NotNil(t, err)
		data.AssertExpectations(t)
	})
}

func TestRegister(t *testing.T) {
	data := mocks.NewAuthDataInterface(t)

	input := users.UserEntity{
		Id:          uint(1),
		TeamId:      1,
		FullName:    "rischi",
		Email:       "rischi@mail",
		Password:    "12345",
		PhoneNumber: "0812345",
		Address:     "pekanbaru",
		Role:        "active",
		Status:      false,
	}
	input2 := users.UserEntity{
		Id:          uint(0),
		TeamId:      0,
		FullName:    "",
		Email:       "",
		Password:    "",
		PhoneNumber: "",
		Address:     "",
		Role:        "",
		Status:      false,
	}
	srv := New(data)

	t.Run("succes", func(t *testing.T) {
		data.On("Store", input).Return(uint(1), nil)
		data.On("SelectById", uint(1)).Return(input, nil)

		err := srv.Register(input)

		assert.NoError(t, err)
		assert.Equal(t, input, err)

		data.AssertExpectations(t)
	})
	t.Run("eror", func(t *testing.T) {
		err := srv.Register(input2)
		assert.NotEmpty(t, err)
		assert.NotNil(t, err)
		data.AssertExpectations(t)
	})
}
