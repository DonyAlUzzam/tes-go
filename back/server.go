package main

import (
	"fmt"
	"io"
	"myapp/config"
	"myapp/model"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
)

type Response struct {
	ErrorCode int         `json:"error_code" form:"error_code"`
	Message   string      `json:"message" form:"message"`
	Data      interface{} `json:"data"`
}

func main() {
	config.ConnectDB()
	route := echo.New()
	route.POST("user/create_user", func(c echo.Context) error {
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
		response := new(Response)
		if user.CreateUser() != nil {
			response.ErrorCode = 10
			response.Message = "Gagal create user"
		} else {
			response.ErrorCode = 0
			response.Message = "Sukses create user"
			response.Data = *user
		}
		return c.JSON(http.StatusOK, response)
	})

	route.PATCH("user/update_user/:email", func(c echo.Context) error {
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
	})

	route.DELETE("user/delete_user/:email", func(c echo.Context) error {
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
	})

	route.GET("user/search_user", func(c echo.Context) error {
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
	})

	route.GET("user/search_user/:email", func(c echo.Context) error {
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
	})

	route.Start(":8000")
}
