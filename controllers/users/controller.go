package users

import (
	"kemahin/app/middlewares"
	"kemahin/businesses/users"
	controller "kemahin/controllers"
	"kemahin/controllers/users/request"
	"kemahin/controllers/users/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userServices users.Service
}

func NewUserController(service users.Service) *UserController {
	return &UserController{
		userServices: service,
	}
}

func (ctrl *UserController) Register(c echo.Context) error {
	req := request.Users{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	data, err := ctrl.userServices.Register(req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(data))
}

func (ctrl *UserController) GetUserRole(id int) string {
	role := ""
	user, err := ctrl.userServices.GetByID(id)
	if err == nil {
		if user.RoleID == 1 {
			role = "mhs"
		}
		if user.RoleID == 2 {
			role = "admin"
		}
		if user.RoleID == 3 {
			role = "organizer"
		}
	}
	return role
}

func (ctrl *UserController) Login(c echo.Context) error {
	req := request.UserLogin{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	token, err := ctrl.userServices.Login(req.NIM, req.Password)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	res := struct {
		Token string `json: "token"`
	}{Token: token}

	return controller.NewSuccessResponse(c, res)
}

func (ctrl *UserController) GetByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))
	user, err := ctrl.userServices.GetByID(id)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controller.NewSuccessResponse(c, user)
}

func (ctrl *UserController) Update(c echo.Context) error {
	req := request.Users{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	id, _ := strconv.Atoi(middlewares.GetUser(c).Id)
	resp, err := ctrl.userServices.Update(id, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controller.NewSuccessResponse(c, resp)
}
