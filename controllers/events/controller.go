package events

import (
	"kemahin/businesses/events"
	controller "kemahin/controllers"
	"kemahin/controllers/events/request"
	"kemahin/controllers/events/response"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type EventController struct {
	eventService events.Service
}

func NewEventController(service events.Service) *EventController {
	return &EventController{
		eventService: service,
	}
}

func (ctrl *EventController) Register(c echo.Context) error {
	req := request.Events{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	data, err := ctrl.eventService.Register(req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controller.NewSuccessResponse(c, response.FromDomain(data))
	// return controller.NewSuccessResponse(c, data)

}

func (ctrl *EventController) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))
	req := request.Events{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	domainReq := req.ToDomain()
	domainReq.Id = id
	resp, err := ctrl.eventService.Update(id, domainReq)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controller.NewSuccessResponse(c, response.FromDomain(resp))
}

func (ctrl *EventController) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := ctrl.eventService.Delete(id)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusNotFound, err)
	}

	return controller.NewSuccessResponse(c, data)
}

func (ctrl *EventController) GetByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.QueryParam("id"))
	event, err := ctrl.eventService.GetByID(id)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, event)
}

func (ctrl *EventController) GetByJudul(c echo.Context) error {
	title := c.QueryParam("title")
	event, err := ctrl.eventService.GetByJudul(title)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(event))
}

func (ctrl *EventController) UpcomingEvent(c echo.Context) error {
	date := time.Now()
	resp, err := ctrl.eventService.UpcomingEvent(date)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	respController := []response.Events{}
	for _, value := range resp {
		respController = append(respController, response.FromDomain(value))
	}

	return controller.NewSuccessResponse(c, respController)
}
