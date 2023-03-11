package delivery

import (
	status "immersiveApp/features/statuses"
	"immersiveApp/utils/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type StatusHandler struct {
	Service status.StatusServiceInterface
}

func New(srv status.StatusServiceInterface) *StatusHandler {
	return &StatusHandler{
		Service: srv,
	}
}

func (t *StatusHandler) GetAll(c echo.Context) error {
	teamEntity, err := t.Service.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFail("error read data"))
	}
	listTeamResponse := ListStatusEntityToStatusResponse(teamEntity)
	return c.JSON(http.StatusOK, helpers.ResponseSuccess("all status", listTeamResponse))
}
