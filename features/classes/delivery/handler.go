package delivery

import (
	"immersiveApp/features/classes"
	"immersiveApp/utils/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ClassHandler struct {
	Service classes.ClassServiceInterface
}

func New(srv classes.ClassServiceInterface) *ClassHandler {
	return &ClassHandler{
		Service: srv,
	}
}

func (t *ClassHandler) GetAll(c echo.Context) error {
	classEntity, err := t.Service.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFail("error read data"))
	}
	listClassResponse := ListClassEntityToClassResponse(classEntity)
	return c.JSON(http.StatusOK, helpers.ResponseSuccess("all class", listClassResponse))
}

func (t *ClassHandler) GetById(c echo.Context) error {
	_id, _ := strconv.Atoi(c.Param("id"))
	id := uint(_id)
	classEntity, err := t.Service.GetById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.ResponseFail("data not found"))
	}
	teamResponse := ClassEntityToClassResponse(classEntity)
	return c.JSON(http.StatusOK, helpers.ResponseSuccess("-", teamResponse))
}

func (t *ClassHandler) Create(c echo.Context) error {
	var formInput ClassRequest
	if err := c.Bind(&formInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error bind data"))
	}

	team, err := t.Service.Create(ClassRequestToClassEntity(&formInput))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusCreated, helpers.ResponseSuccess("Create Data Success", ClassEntityToClassResponse(team)))
}

func (t *ClassHandler) Update(c echo.Context) error {
	var formInput ClassRequest
	if err := c.Bind(&formInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error bind data"))
	}

	_id, _ := strconv.Atoi(c.Param("id"))
	id := uint(_id)

	team, err := t.Service.Update(ClassRequestToClassEntity(&formInput), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("Update Data Success", ClassEntityToClassResponse(team)))
}

func (t *ClassHandler) Delete(c echo.Context) error {
	_id, _ := strconv.Atoi(c.Param("id"))
	id := uint(_id)

	if err := t.Service.Delete(id); err != nil {
		return c.JSON(http.StatusNotFound, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("Delete Data Success", nil))
}
