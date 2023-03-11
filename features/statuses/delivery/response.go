package delivery

import (
	status "immersiveApp/features/statuses"
)

type StatusResp struct {
	ID          uint
	Name        string
	Description string
}

type StatusResponse struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func StatusRequestToStatusEntity(statusRequest *StatusRequest) status.StatusEntity {
	return status.StatusEntity{
		Name:        statusRequest.Name,
		Description: statusRequest.Description,
	}
}

func StatusEntityToSTatusResponse(status status.StatusEntity) StatusResp {
	return StatusResp{
		ID:          status.Id,
		Name:        status.Name,
		Description: status.Description,
	}
}

func ListStatusEntityToStatusResponse(statusEntity []status.StatusEntity) []StatusResp {
	var dataResponses []StatusResp
	for _, v := range statusEntity {
		dataResponses = append(dataResponses, StatusEntityToSTatusResponse(v))
	}
	return dataResponses
}
