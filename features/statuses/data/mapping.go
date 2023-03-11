package data

import (
	"immersiveApp/features/statuses"
)

func StatusEntityToStatus(statusEntity statuses.StatusEntity) Status {
	return Status{
		Name:        statusEntity.Name,
		Description: statusEntity.Description,
	}
}

func StatusToStatusEntity(status Status) statuses.StatusEntity {
	return statuses.StatusEntity{
		Id:        status.ID,
		Name:      status.Name,
		CreatedAt: status.CreatedAt,
		UpdatedAt: status.UpdatedAt,
	}
}

func ListStatusToStatusEntity(status []Status) []statuses.StatusEntity {
	var statusEntity []statuses.StatusEntity
	for _, v := range status {
		statusEntity = append(statusEntity, StatusToStatusEntity(v))
	}
	return statusEntity
}
