package delivery

import (
	"clean-arch/features/users"
	"clean-arch/utils/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	Service users.UserServiceInterface
}

func New(s users.UserServiceInterface) *UserHandler {
	return &UserHandler{
		Service: s,
	}
}

func (u *UserHandler) GetAll(c echo.Context) error {
	users, err := u.Service.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("-", users))
}

func (u *UserHandler) GetById(c echo.Context) error {
	_id, _ := strconv.Atoi(c.Param("id"))
	id := uint(_id)

	user, err := u.Service.GetById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFail(err.Error()))
	}

	data := UserEntityToUserResponse(user)

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("-", data))
}

func (u *UserHandler) Create(c echo.Context) error {
	var formInput UserRequest
	if err := c.Bind(&formInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error bind data"))
	}

	user := UserRequestToUserEntity(formInput)
	if err := u.Service.Create(user); err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusCreated, helpers.ResponseSuccess("Register Success", UserEntityToUserResponse(user)))
}

func (u *UserHandler) Update(c echo.Context) error {
	var formInput UserRequest
	if err := c.Bind(&formInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error bind data"))
	}

	_id, _ := strconv.Atoi(c.Param("id"))
	id := uint(_id)

	user := UserRequestToUserEntity(formInput)
	if err := u.Service.Update(user, id); err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("berhasil update data user", UserEntityToUserResponse(user)))
}

func (u *UserHandler) Delete(c echo.Context) error {
	_id, _ := strconv.Atoi(c.Param("id"))
	id := uint(_id)

	if err := u.Service.Delete(id); err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("berhasil delete data", nil))
}
