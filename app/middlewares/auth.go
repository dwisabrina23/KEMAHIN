package middlewares

import (
	controller "kemahin/controllers"
	"net/http"

	"time"
	// "strings"
	// "github.com/brianvoe/sjwt"
	"github.com/golang-jwt/jwt"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JwtCustomClaims struct {
	ID int `json:"id"`
	// Role int `json:"role"`
	jwt.StandardClaims
}

type ConfigJWT struct {
	SecretJWT       string
	ExpiresDuration int
}

func (jwtConf *ConfigJWT) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: []byte(jwtConf.SecretJWT),
		ErrorHandlerWithContext: middleware.JWTErrorHandlerWithContext(func(e error, c echo.Context) error {
			return controller.NewErrorResponse(c, http.StatusForbidden, e)
		}),
	}
}

// GenerateToken jwt ...
func (jwtConf *ConfigJWT) GenerateToken(userID int) string {
	claims := JwtCustomClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(jwtConf.ExpiresDuration))).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, _ := token.SignedString([]byte(jwtConf.SecretJWT))

	return jwtToken
}

//get user by jwt
func GetUser(c echo.Context) *JwtCustomClaims {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)
	return claims
}

// func AdminRoleValidation(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(e echo.Context) error {
// 		role, err := ExtractJWTPayloadRole(e)
// 		if err != nil {
// 			return echo.ErrUnauthorized
// 		}
// 		if role == "admin" {
// 			return next(e)
// 		}
// 		return echo.ErrUnauthorized
// 	}
// }

// func UserRoleValidation(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(e echo.Context) error {
// 		role, err := ExtractJWTPayloadRole(e)
// 		if err != nil {
// 			return echo.ErrUnauthorized
// 		}
// 		if role == "user" {
// 			return next(e)
// 		}
// 		return echo.ErrUnauthorized
// 	}
// }

// func ExtractJWTPayloadRole(c echo.Context) (string, error) {
// 	header := c.Request().Header.Clone().Get("Authorization")
// 	token := strings.Split(header, "Bearer ")[1]
// 	claims, err := sjwt.Parse(token)
// 	if err != nil {
// 		return "", err
// 	}
// 	return claims["role"].(string), nil
// }

// func ExtractJWTPayloadUserId(c echo.Context) (float64, error) {
// 	header := c.Request().Header.Clone().Get("Authorization")
// 	token := strings.Split(header, "Bearer ")[1]
// 	claims, _ := sjwt.Parse(token)
// 	userId := claims["user_id"].(float64)
// 	return userId, nil
// }
