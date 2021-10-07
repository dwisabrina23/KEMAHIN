package orders

import (
	"kemahin/app/middlewares"
	"kemahin/businesses/orders"
	controller "kemahin/controllers"
	"kemahin/controllers/orders/request"
	"kemahin/controllers/orders/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type OrderController struct {
	orderService orders.Service
}

func NewOrderController(service orders.Service) *OrderController {
	return &OrderController{
		orderService: service,
	}
}

func (ctrl *OrderController) Create(c echo.Context) error {
	req := request.Orders{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	userID := middlewares.GetUser(c).ID

	data, err := ctrl.orderService.Create(userID, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controller.NewSuccessResponse(c, data)
}

func (ctrl *OrderController) GetByUserID(c echo.Context) error {
	userID := middlewares.GetUser(c).ID
	data, err := ctrl.orderService.GetByUserID(userID)

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusNotFound, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomainArray(data))
}

func (ctrl *OrderController) ValidateOrder(c echo.Context) error {
	orderID, _ := strconv.Atoi(c.Param("order_id"))
	resp, err := ctrl.orderService.ValidateOrder(orderID)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusNotFound, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(resp))
}
