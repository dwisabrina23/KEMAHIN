package organizers

import (
	"kemahin/businesses/organizers"
	controller "kemahin/controllers"
	"kemahin/controllers/organizer/request"
	"kemahin/controllers/organizer/response"
	"strconv"

	"net/http"

	"github.com/labstack/echo/v4"
)

type OrgController struct {
	orgServices organizers.Service
}

func NewOrgController(service organizers.Service) *OrgController {
	return &OrgController{
		orgServices: service,
	}
}

func (ctrl *OrgController) Register(c echo.Context) error {
	req := request.Organizers{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	data, err := ctrl.orgServices.Register(req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(data))
}

func (ctrl *OrgController) Login(c echo.Context) error {
	req := request.OrgLogin{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	token, err := ctrl.orgServices.Login(req.Username, req.Password)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	res := struct {
		Token string `json: "token"`
	}{Token: token}

	return controller.NewSuccessResponse(c, res)
}

func (ctrl *OrgController) GetByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))
	user, err := ctrl.orgServices.GetByID(id)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controller.NewSuccessResponse(c, user)
}
