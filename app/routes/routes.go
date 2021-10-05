package routes

import (
	middlewareApp "kemahin/app/middlewares"
	"kemahin/controllers/users"

	"errors"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	controller "kemahin/controllers"
	"net/http"
)

type ControllerList struct {
	JWTMiddleware  middleware.JWTConfig
	UserController users.UserController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	middlewareApp.Log(e)
	e.Pre(middleware.RemoveTrailingSlash())
	users := e.Group("user")
	users.POST("/login", cl.UserController.Login)
	users.POST("/register", cl.UserController.Register)
	users.GET("", cl.UserController.GetByID)
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
