package routes

import (
	middlewareApp "kemahin/app/middlewares"
	"kemahin/controllers/events"
	"kemahin/controllers/users"

	"errors"
	controller "kemahin/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware   middleware.JWTConfig
	UserController  users.UserController
	EventController events.EventController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	// middlewareApp.Log(e)
	// e.Pre(middleware.RemoveTrailingSlash))
	users := e.Group("user")
	users.POST("/login", cl.UserController.Login)
	users.POST("/register", cl.UserController.Register)
	users.GET("/:id", cl.UserController.GetByID)
	users.PUT("/update/:id", cl.UserController.Update, middleware.JWTWithConfig(cl.JWTMiddleware))

	events := e.Group("events")
	events.POST("/register", cl.EventController.Register)
	events.PUT("/:id", cl.EventController.Update)
	events.DELETE(":/id", cl.EventController.Delete)
	events.GET("/:id", cl.EventController.GetByID)
	events.GET("/:judul", cl.EventController.GetByJudul)
	events.GET("/upcoming", cl.EventController.UpcomingEvent)
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
