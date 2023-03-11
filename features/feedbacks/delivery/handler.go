package delivery

import (
	"immersiveApp/app/middlewares"
	"immersiveApp/features/feedbacks"
	"immersiveApp/utils/helpers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type FeedbackHandler struct {
	Service feedbacks.FeedbackServiceInterface
	
}

func New(srv feedbacks.FeedbackServiceInterface) *FeedbackHandler {
	return &FeedbackHandler{
		Service: srv,
	}
}

// not used
func (f *FeedbackHandler) GetAll(c echo.Context) error {
	feedbackEntity, err := f.Service.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFail("error read data"))
	}
	listFeedbackResponse := ListFeedbackEntityToFeedbackResponse(feedbackEntity)
	return c.JSON(http.StatusOK, helpers.ResponseSuccess("all feedback", listFeedbackResponse))
}

// not used
func (f *FeedbackHandler) GetById(c echo.Context) error {
	_id, _ := strconv.Atoi(c.Param("id"))
	id := uint(_id)
	feedbackEntity, err := f.Service.GetById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.ResponseFail("data not found"))
	}
	feedbackResponse := FeedbackEntityToFeedbackResponse(feedbackEntity)
	return c.JSON(http.StatusOK, helpers.ResponseSuccess("-", feedbackResponse))
}

func (f *FeedbackHandler) Create(c echo.Context) error {
	var formInput FeedbackRequest
	if err := c.Bind(&formInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error bind data"))
	}

	claim := middlewares.ClaimsToken(c)
	user_id := claim.Id
	formInput.UserId = uint(user_id)

	feedback, err := f.Service.Create(FeedbackRequestToFeedbackEntity(&formInput))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusCreated, helpers.ResponseSuccess("Create Data Success", FeedbackEntityToFeedbackResponse(feedback)))
}

func (f *FeedbackHandler) Update(c echo.Context) error {
	var formInput FeedbackRequest
	if err := c.Bind(&formInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error bind data"))
	}

	_id, _ := strconv.Atoi(c.Param("id"))
	id := uint(_id)

	feedback, err := f.Service.Update(FeedbackRequestToFeedbackEntity(&formInput), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("Update Data Success", FeedbackEntityToFeedbackResponse(feedback)))
}

func (f *FeedbackHandler) Delete(c echo.Context) error {
	_id, _ := strconv.Atoi(c.Param("id"))
	id := uint(_id)

	if err := f.Service.Delete(id); err != nil {
		return c.JSON(http.StatusNotFound, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("Delete Data Success", nil))
}
