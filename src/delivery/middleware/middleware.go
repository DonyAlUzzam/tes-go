package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"

	"os"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Nama  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var (
	accessTokenCookieName  = os.Getenv("ACCESSTOKENCOOKIE")
	refreshTokenCookieName = os.Getenv("REFRESHTOKEN")
	jwtSecretKey           = os.Getenv("JWTSECRETKEY")
	jwtRefreshSecretKey    = os.Getenv("JWTREFRESHKEY")
)

func GetJWTSecret() string {
	return jwtSecretKey
}

type AccessToken struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GenerateToken(accessToken AccessToken) string {
	expirationTime := time.Now().Add(1 * time.Hour)

	claims := &Claims{
		Nama:  accessToken.Name,
		Email: accessToken.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, _ := token.SignedString([]byte(GetJWTSecret()))

	return tokenString
}

func JWTVerify(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		request := c.Request()
		tokenHeader := request.Header.Get("token")
		tokenArr := strings.Split(tokenHeader, " ")
		tes := len(tokenArr)
		if tes == 1 {
			return echo.NewHTTPError(http.StatusForbidden, "Token Not Valid")
		}

		if tokenArr[0] != "Bearer" {
			return echo.NewHTTPError(http.StatusForbidden, "Invalid Token")
		}

		tokenString := tokenArr[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(GetJWTSecret()), nil
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusForbidden, "UnAuthorized")
		}
		_, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return echo.NewHTTPError(http.StatusForbidden, "UnAuthorized")
		}
		return next(c)
	}
}
