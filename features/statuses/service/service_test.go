package service

import (
	"immersiveApp/features/statuses"
	"immersiveApp/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestStatus(t *testing.T) {
	repo := mocks.NewStatusDataInterface(t)
	expectedData := []statuses.StatusEntity{
		{
			Id:   1,
			Name: "Status 1",
		},
		{
			Id:   2,
			Name: "Status 2",
		},
	}

	t.Run("Sukses lihat status", func(t *testing.T) {
		repo.On("SelectAll", mock.Anything).Return(expectedData, nil).Once()
		srv := New(repo)
		result, err := srv.GetAll()
		assert.Nil(t, err)
		assert.Equal(t, expectedData, result)
		repo.AssertCalled(t, "SelectAll")
		repo.AssertExpectations(t)
	})
}
