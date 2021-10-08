package tickets

import (
	"kemahin/businesses/tickets"
	controller "kemahin/controllers"
	"kemahin/controllers/tickets/request"
	"kemahin/controllers/tickets/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TicketController struct {
	ticketService tickets.Service
}

func NewTicketController(service tickets.Service) *TicketController {
	return &TicketController{
		ticketService: service,
	}
}

func (ctrl *TicketController) Create(c echo.Context) error {
	req := request.Tickets{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	data, err := ctrl.ticketService.Create(req.OrderID, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, data)

}
func (ctrl *TicketController) GetByUserId(c echo.Context) error {
	userID, _ := strconv.Atoi(c.Param("user_id"))
	data, err := ctrl.ticketService.GetByUserId(userID)

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusNotFound, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomainArray(data))
}
