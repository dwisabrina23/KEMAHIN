package controller

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Meta struct {
		Status   int      `json:"status"`
		Message  string   `json:"message"`
		Messages []string `json:"messages,omitempty"`
	} `json:"meta"`
	Data interface{} `json:"data"`
}

//update response
type UpdateResponse struct {
	Meta struct {
		Status  int    `json:"status"`
		Message string `json:"messsage"`
	}
}

func NewSuccessResponse(c echo.Context, param interface{}) error {
	response := BaseResponse{}
	response.Meta.Status = http.StatusOK
	response.Meta.Message = "Success"
	response.Data = param

	return c.JSON(http.StatusOK, response)
}

//update success response
func UpdateSuccesResponse(c echo.Context, message string) error {
	response := UpdateResponse{}
	response.Meta.Status = http.StatusOK
	response.Meta.Message = message
	return c.JSON(http.StatusOK, response)
}

func NewErrorResponse(c echo.Context, status int, err error) error {
	response := BaseResponse{}
	response.Meta.Status = status
	response.Meta.Message = "Something not right"
	response.Meta.Messages = []string{err.Error()}

	return c.JSON(status, response)
}
