package service

import (
	"errors"
	tims "immersiveApp/features/teams"
	"immersiveApp/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	repo := mocks.NewTeamDataInterface(t)
	input := tims.TeamEntity{
		Id:   uint(1),
		Name: "",
	}
	srv := New(repo)

	mockTeamEntity := tims.TeamEntity{
		Name: "Team A",
	}
	t.Run("succes", func(t *testing.T) {
		repo.On("Store", mockTeamEntity).Return(uint(1), nil)
		repo.On("SelectById", uint(1)).Return(mockTeamEntity, nil)

		created, err := srv.Create(mockTeamEntity)

		assert.NoError(t, err)
		assert.Equal(t, mockTeamEntity, created)

		repo.AssertExpectations(t)
	})

	t.Run("eror", func(t *testing.T) {
		_, err := srv.Create(input)
		assert.NotEmpty(t, err)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestGetAll(t *testing.T) {
	repo := mocks.NewTeamDataInterface(t)
	mockTeam := tims.TeamEntity{
		Id:   0,
		Name: "team1",
	}

	mockListTeam := make([]tims.TeamEntity, 0)
	mockListTeam = append(mockListTeam, mockTeam)

	repo.On("SelectAll").Return(mockListTeam, nil)

	srv := New(repo)

	result, err := srv.GetAll()

	repo.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Len(t, result, len(mockListTeam))
	assert.Equal(t, mockListTeam, result)

}
func TestGetById(t *testing.T) {
	repo := mocks.NewTeamDataInterface(t)
	mockTeam := tims.TeamEntity{
		Id:   1,
		Name: "team1",
	}

	repo.On("SelectById", mockTeam.Id).Return(mockTeam, nil)

	srv := New(repo)

	result, err := srv.GetById(mockTeam.Id)
	repo.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, mockTeam, result)
}

func TestDelete(t *testing.T) {
	repo := mocks.NewTeamDataInterface(t)

	mockTeamEntity := tims.TeamEntity{
		Name: "Team A",
	}
	t.Run("Sukses Delete", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(mockTeamEntity, nil).Once()
		srv := New(repo)
		err := srv.Delete(1)
		assert.Nil(t, err)
		assert.NotEmpty(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Gagal Delete", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(mockTeamEntity, errors.New("error")).Once()
		srv := New(repo)
		err := srv.Delete(1)
		assert.NotNil(t, err)
		assert.NotEmpty(t, err)
		repo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	repo := mocks.NewTeamDataInterface(t)
	expected := tims.TeamEntity{
		Id:        1,
		Name:      "team A",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	input := tims.TeamEntity{
		Id:        1,
		Name:      "team B",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	mockTeamEntity := tims.TeamEntity{
		Name: "Team A",
	}
	t.Run("Update success", func(t *testing.T) {
		repo.On("Update", expected, 1).Return(input, nil).Once()
		srv := New(repo)

		res, err := srv.Update(input, 1)
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})

	t.Run("eror", func(t *testing.T) {
		srv := New(repo)
		_, err := srv.Update(mockTeamEntity, 0)
		assert.NotEmpty(t, err)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Update Fail", func(t *testing.T) {
		srv := New(repo)
		res, err := srv.Update(input, 0)
		assert.Empty(t, res)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}
