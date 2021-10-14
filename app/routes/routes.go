package routes

import (
	middlewareApp "kemahin/app/middlewares"
	"kemahin/controllers/events"
	"kemahin/controllers/orders"
	"kemahin/controllers/organizer"
	"kemahin/controllers/tickets"
	"kemahin/controllers/users"

	"errors"
	controller "kemahin/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware       middleware.JWTConfig
	UserController      users.UserController
	EventController     events.EventController
	OrganizerController organizers.OrgController
	OrdersController    orders.OrderController
	TicketController    tickets.TicketController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	middlewareApp.Log(e)
	e.Pre(middleware.RemoveTrailingSlash())
	users := e.Group("user")
	users.POST("/login", cl.UserController.Login)
	users.POST("/register", cl.UserController.Register)
	users.GET("/:id", cl.UserController.GetByID)
	users.PUT("/update/:id", cl.UserController.Update, middleware.JWTWithConfig(cl.JWTMiddleware))
	users.PUT("/updates", cl.UserController.Update, middleware.JWTWithConfig(cl.JWTMiddleware))
	users.PUT("/update", cl.UserController.Update)

	events := e.Group("events")
	events.POST("/register", cl.EventController.Register)
	events.PUT("/:id", cl.EventController.Update)
	events.DELETE(":/id", cl.EventController.Delete)
	events.GET("/:id", cl.EventController.GetByID)
	events.GET("/:judul", cl.EventController.GetByJudul)
	events.GET("/upcoming", cl.EventController.UpcomingEvent)

	org := e.Group("organizer")
	org.POST("/register", cl.OrganizerController.Register)
	org.POST("/login", cl.OrganizerController.Login)

	orders := e.Group("orders")
	orders.POST("/create", cl.OrdersController.Create)
	orders.GET("/:id", cl.UserController.GetByID)
	orders.PUT("/validate/:id", cl.OrdersController.ValidateOrder)

	tickets := e.Group("tickets")
	tickets.POST("/create", cl.TicketController.Create)
	tickets.GET("/:user_id", cl.TicketController.GetByUserId)

	sends := e.Group("sends")
	sends.POST("/:id", cl.TicketController.GetByUserId)
}

func RoleValidation(role string, userController users.UserController) echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := middlewareApp.GetUser(c)
			userRole := userController.GetUserRole(claims.ID)
			if userRole == role {
				return hf(c)
			} else {
				return controller.NewErrorResponse(c, http.StatusForbidden, errors.New("Unauthorized"))
			}
		}
	}
}
