package service

import (
	"immersiveApp/features/classes"
	"immersiveApp/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAll(t *testing.T) {
	repo := mocks.NewClassDataInterface(t)
	mock := classes.ClassEntity{
		Id:           0,
		ClassName:    "",
		PicUserId:    0,
		DateStart:    "",
		DateGraduate: "",
	}

	mockList := make([]classes.ClassEntity, 0)
	mockList = append(mockList, mock)

	repo.On("SelectAll").Return(mockList, nil)

	srv := New(repo)

	result, err := srv.GetAll()

	repo.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Len(t, result, len(mockList))
	assert.Equal(t, mockList, result)
}
func TestGetById(t *testing.T) {
	repo := mocks.NewClassDataInterface(t)
	mock := classes.ClassEntity{
		Id:           1,
		ClassName:    "BE 15",
		PicUserId:    1,
		DateStart:    "2023-03-03",
		DateGraduate: "2023-03-04",
	}

	repo.On("SelectById", mock.Id).Return(mock, nil)

	srv := New(repo)

	result, err := srv.GetById(mock.Id)
	repo.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, mock, result)
}

func TestCreate(t *testing.T) {
	repo := mocks.NewClassDataInterface(t)
	input := classes.ClassEntity{
		Id:           uint(1),
		ClassName:    "",
		PicUserId:    1,
		DateStart:    "",
		DateGraduate: "",
	}
	srv := New(repo)

	mock := classes.ClassEntity{
		ClassName:    "BE 15",
		PicUserId:    1,
		DateStart:    "2023-03-03",
		DateGraduate: "2023-03-04",
	}
	t.Run("succes", func(t *testing.T) {
		repo.On("Store", mock).Return(uint(1), nil)
		repo.On("SelectById", uint(1)).Return(mock, nil)

		created, err := srv.Create(mock)

		assert.NoError(t, err)
		assert.Equal(t, mock, created)

		repo.AssertExpectations(t)
	})

	t.Run("eror", func(t *testing.T) {
		_, err := srv.Create(input)
		assert.NotEmpty(t, err)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	repo := mocks.NewClassDataInterface(t)

	mockTeamEntity := classes.ClassEntity{
		Id:           1,
		ClassName:    "BE 15",
		PicUserId:    1,
		DateStart:    "2023-03-03",
		DateGraduate: "2023-03-04",
	}
	t.Run("Sukses Delete", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(mockTeamEntity, nil).Once()
		repo.On("SelectById", uint(1)).Return(mockTeamEntity, nil)
		srv := New(repo)
		err := srv.Delete(1)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("Gagal Delete", func(t *testing.T) {
		repo.On("SelectById", uint(1)).Return(mockTeamEntity, nil)
		srv := New(repo)
		err := srv.Delete(1)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	repo := mocks.NewClassDataInterface(t)
	expected := classes.ClassEntity{
		Id:           1,
		ClassName:    "BE 15",
		PicUserId:    1,
		DateStart:    "2023-03-03",
		DateGraduate: "2023-03-04",
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
	}

	input := classes.ClassEntity{
		Id:           1,
		ClassName:    "BE 15",
		PicUserId:    1,
		DateStart:    "2023-03-03",
		DateGraduate: "2023-03-04",
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
	}

	mockTeamEntity := classes.ClassEntity{
		ClassName:    "BE 15",
		PicUserId:    1,
		DateStart:    "2023-03-03",
		DateGraduate: "2023-03-04",
	}

	t.Run("Update success", func(t *testing.T) {
		repo.On("Update", expected, 1).Return(input, nil).Once()
		srv := New(repo)

		res, err := srv.Update(input, 1)
		assert.Nil(t, err)
		assert.NotEmpty(t, res)
		repo.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
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
