package delivery

import (
	"clean-arch/features/teams"
	"clean-arch/utils/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TeamHandler struct {
	Service teams.TeamServiceInterface
}

func New(srv teams.TeamServiceInterface) *TeamHandler {
	return &TeamHandler{
		Service: srv,
	}
}

func (t *TeamHandler) GetAll(c echo.Context) error {
	teamEntity, err := t.Service.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFail("error read data"))
	}
	listTeamResponse := ListTeamEntityToTeamResponse(teamEntity)
	return c.JSON(http.StatusOK, helpers.ResponseSuccess("all teams", listTeamResponse))
}

func (t *TeamHandler) GetById(c echo.Context) error {
	_id, _ := strconv.Atoi(c.Param("id"))
	id := uint(_id)
	teamEntity, err := t.Service.GetById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFail("error read data"))
	}
	teamResponse := TeamEntityToTeamResponse(teamEntity)
	return c.JSON(http.StatusOK, helpers.ResponseSuccess("all teams", teamResponse))
}

func (t *TeamHandler) Create(c echo.Context) error {
	var formInput TeamRequest
	if err := c.Bind(&formInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error bind data"))
	}

	team := TeamRequestToTeamEntity(formInput)
	if err := t.Service.Create(team); err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusCreated, helpers.ResponseSuccess("Register Success", TeamEntityToTeamResponse(team)))
}

func (t *TeamHandler) Update(c echo.Context) error {
	var formInput TeamRequest
	if err := c.Bind(&formInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error bind data"))
	}

	_id, _ := strconv.Atoi(c.Param("id"))
	id := uint(_id)

	team := TeamRequestToTeamEntity(formInput)
	if err := t.Service.Update(team, id); err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("berhasil update data user", TeamEntityToTeamResponse(team)))
}

func (t *TeamHandler) Delete(c echo.Context) error {
	_id, _ := strconv.Atoi(c.Param("id"))
	id := uint(_id)

	if err := t.Service.Delete(id); err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("berhasil delete data", nil))
}
