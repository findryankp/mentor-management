package service

import (
	"errors"
	"immersiveApp/features/feedbacks"
	"immersiveApp/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	repo := mocks.NewFeedbackDataInterface(t)
	input := feedbacks.FeedbackEntity{
		Id:        uint(1),
		Notes:     "ini notes",
		Proof:     "ini bukti",
		UserId:    1,
		MenteeId:  1,
		StatusId:  1,
		Approved:  false,

	}
	srv := New(repo)

	mockinput := feedbacks.FeedbackEntity{
		Id:        uint(2),
		Notes:     "",
		Proof:     "",
		UserId:    0,
		MenteeId:  0,
		StatusId:  0,
		Approved:  false,
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
	repo := mocks.NewFeedbackDataInterface(t)
	input := feedbacks.FeedbackEntity{
		Id:        uint(1),
		Notes:     "",
		Proof:     "",
		UserId:    0,
		MenteeId:  0,
		StatusId:  0,
		Approved:  false,
	}
	srv := New(repo)

	mockList := make([]feedbacks.FeedbackEntity, 0)
	mockList = append(mockList, input)

	repo.On("SelectAll").Return(mockList, nil)

	result, err := srv.GetAll()

	repo.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Len(t, result, len(mockList))
	assert.Equal(t, mockList, result)

}
func TestGetById(t *testing.T) {
	repo := mocks.NewFeedbackDataInterface(t)
	input := feedbacks.FeedbackEntity{
		Id:        uint(1),
		Notes:     "ini notes",
		Proof:     "ini bukti",
		UserId:    1,
		MenteeId:  1,
		StatusId:  1,
		Approved:  true,
	}

	repo.On("SelectById", input.Id).Return(input, nil)

	srv := New(repo)

	result, err := srv.GetById(input.Id)
	repo.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, input, result)
}

func TestDelete(t *testing.T) {
	repo := mocks.NewFeedbackDataInterface(t)
	input := feedbacks.FeedbackEntity{
		Id:        uint(1),
		Notes:     "ini notes",
		Proof:     "ini bukti",
		UserId:    1,
		MenteeId:  3,
		StatusId:  1,
		Approved:  true,

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
		repo.On("Delete", mock.Anything).Return(input, errors.New("error")).Once()
		srv := New(repo)
		err := srv.Delete(1)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	repo := mocks.NewFeedbackDataInterface(t)

	input := feedbacks.FeedbackEntity{
		Id:        uint(1),
		Notes:     "ini notes",
		Proof:     "ini bukti",
		UserId:    1,
		MenteeId:  3,
		StatusId:  1,
		Approved:  true,
	}
	updatedData := feedbacks.FeedbackEntity{
		Id:        uint(1),
		Notes:     "ini notes",
		Proof:     "ini bukti",
		UserId:    1,
		MenteeId:  3,
		StatusId:  1,
		Approved:  true,
	}
	t.Run("sukses update data", func(t *testing.T) {
		repo.On("Update", uint(1), input).Return(updatedData, nil).Once()

		s := &feedbackService{Data: repo}

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

