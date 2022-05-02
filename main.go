package main

import (
	"fmt"
	"io"
	"log"
	"myapp/auth"
	"myapp/config"
	"myapp/model"
	"myapp/utils"
	"net/http"
	"os"
	"strings"

	// "github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
	user := new(model.Users)
	c.Bind(user)
	contentType := c.Request().Header.Get("Content-Type")
	if contentType == "application/json" {
		fmt.Println("request dari json")
	} else if strings.Contains(contentType, "multipart/form-data") || contentType == "application/x-www-form-urlencoded" {
		file, err := c.FormFile("ktp")
		if err != nil {
			fmt.Println("ktp kosong")
		} else {
			src, err := file.Open()
			if err != nil {
				return err
			}
			defer src.Close()
			dst, err := os.Create(file.Filename)
			if err != nil {
				return err
			}
			defer dst.Close()
			if _, err = io.Copy(dst, src); err != nil {
				return err
			}

			user.Ktp = file.Filename
			fmt.Println("ada file, akan disimpan")
		}

	}
	response := new(utils.Response)
	if user.CreateUser() != nil {
		response.ErrorCode = 10
		response.Message = "Gagal create user"
	} else {
		response.ErrorCode = 0
		response.Message = "Sukses create user"
		response.Data = *user
	}
	return c.JSON(http.StatusOK, response)
}

func UpdateUser(c echo.Context) error {
	user := new(model.Users)
	c.Bind(user)
	response := new(Response)
	if user.UpdateUser(c.Param("email")) != nil {
		response.ErrorCode = 10
		response.Message = "Gagal updatre user"
	} else {
		response.ErrorCode = 0
		response.Message = "Sukses update user"
		response.Data = *user
	}
	return c.JSON(http.StatusOK, response)
}

func DeleteUser(c echo.Context) error {
	user, _ := model.GetOneByEmail(c.Param("email"))
	response := new(Response)

	if user.DeleteUser() != nil {
		response.ErrorCode = 10
		response.Message = "Gagal menghapus data"
	} else {
		response.ErrorCode = 0
		response.Message = "sukses menghapus data"
	}
	return c.JSON(http.StatusOK, response)
}

func GetUserAll(c echo.Context) error {
	response := new(Response)
	users, err := model.GetAll(c.QueryParam("keywords"))

	if err != nil {
		response.ErrorCode = 10
		response.Message = "Gagal melihat data"
	} else {
		response.ErrorCode = 0
		response.Message = "berhasil melihat data"
		response.Data = users
	}
	return c.JSON(http.StatusOK, response)
}

func GetUserByEmail(c echo.Context) error {
	response := new(Response)
	user, err := model.GetOneByEmail(c.Param("email"))
	if err != nil {
		response.ErrorCode = 10
		response.Message = "Gagal melihat data"
	} else {
		response.ErrorCode = 0
		response.Message = "berhasil melihat data"
		response.Data = user
	}

	return c.JSON(http.StatusOK, response)
}

func Login(c echo.Context) error {
	response := new(ResponseLogin)
	form := auth.LoginRequest{}
	c.Bind(&form)

	config.ConnectDB()
	user, err := model.GetOneByEmail(form.Email)
	if user.Password != form.Password {
		return err
	}
	token := auth.GenerateToken(auth.AccessToken{
		Name:  user.Nama,
		Email: user.Email,
	})
	response.ErrorCode = 0
	response.Message = "berhasil login"
	response.Token = token

	return c.JSON(http.StatusOK, response)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDB()
	config.ConnectSolr()
	route := echo.New()
	v1 := route.Group("api/v1/users")
	{
		v1.POST("/", CreateUser)
		v1.POST("/login", Login)
		v1.PATCH("/:email", UpdateUser)
		v1.GET("/", GetUserAll, auth.JWTVerify)
		v1.GET("/:email", GetUserByEmail)
		v1.DELETE("/:email", DeleteUser)
	}

	v2 := route.Group("api/v2/users")
	{
		v2.POST("/", CreateUser)
		v2.PATCH("/:email", UpdateUser)
		v2.GET("/", GetUserAll)
		v2.GET("/:email", GetUserByEmail)
		v2.DELETE("/:email", DeleteUser)
	}

	route.Start(":8000")
}
