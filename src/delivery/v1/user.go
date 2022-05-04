package v1

import (
	"myapp/src/delivery/middleware"
	"myapp/src/usecase"
	users "myapp/src/usecase/user"
	"myapp/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserDelivery struct {
	userUc usecase.UserUsecase
}

func NewUserDelivery(userUc usecase.UserUsecase) *UserDelivery {
	return &UserDelivery{
		userUc: userUc,
	}
}

func (c *UserDelivery) Mount(group *echo.Group) {
	group.GET("", c.GetAllDataUsers)
	group.DELETE("/delete", c.DeleteUserByEmail, middleware.JWTVerify)
	group.PATCH("/update", c.UpdateDataUsers, middleware.JWTVerify)
	group.GET("/detail", c.GetOneByEmail, middleware.JWTVerify)
	group.POST("/login", c.Login)
	group.POST("/create", c.CreateUser, middleware.JWTVerify)
}

func (c *UserDelivery) GetOneByEmail(e echo.Context) error {

	ctx := e.Request().Context()
	form := new(users.GetOneByEmailRequest)

	if err := e.Bind(form); err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	user, err := c.userUc.GetOneByEmail(ctx, form)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	response := new(utils.Response)
	response.Message = "berhasil melihat data"
	response.Data = user

	return e.JSON(http.StatusOK, response)
}

func (c *UserDelivery) GetAllDataUsers(e echo.Context) error {

	ctx := e.Request().Context()
	form := new(users.GetAllDataUsersRequest)

	if err := e.Bind(form); err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	user, err := c.userUc.GetAllDataUsers(ctx, form)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	response := new(utils.Response)
	response.Message = "berhasil melihat data"
	response.Data = user

	return e.JSON(http.StatusOK, response)
}

func (c *UserDelivery) UpdateDataUsers(e echo.Context) error {

	ctx := e.Request().Context()
	form := new(users.UpdateDataByEmailRequest)

	if err := e.Bind(form); err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	user, err := c.userUc.UpdateDataByEmail(ctx, form)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	response := new(utils.Response)
	response.Message = "berhasil update data"
	response.Data = user

	return e.JSON(http.StatusOK, response)
}

func (c *UserDelivery) DeleteUserByEmail(e echo.Context) error {

	ctx := e.Request().Context()
	form := new(users.GetOneByEmailRequest)

	if err := e.Bind(form); err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	response := new(utils.Response)

	_, err := c.userUc.GetOneByEmail(ctx, form)
	if err != nil {
		response.ErrorCode = http.StatusBadRequest
		response.Message = "Data Not Found!"
		return e.JSON(http.StatusBadRequest, response)
	} else {
		_, err := c.userUc.DeleteUserByEmail(ctx, form)
		if err != nil {
			return e.JSON(http.StatusBadRequest, err)
		}

		response.ErrorCode = http.StatusOK
		response.Message = "berhasil delete data"
		response.Data = form
	}

	return e.JSON(http.StatusOK, response)
}

func (c *UserDelivery) Login(e echo.Context) error {
	ctx := e.Request().Context()
	// form := users.LoginRequest{}
	form := new(users.LoginRequest)

	if err := e.Bind(form); err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	response := new(utils.ResponseLogin)

	user, err := c.userUc.Login(ctx, form)
	if err != nil {
		response.Message = "Login Gagal"
		return e.JSON(http.StatusUnauthorized, response)
	}

	if user.Password != form.Password {
		response.Message = "Login Gagal"
		return e.JSON(http.StatusUnauthorized, response)
	}

	token := middleware.GenerateToken(middleware.AccessToken{
		Name:  user.Nama,
		Email: user.Email,
	})
	response.ErrorCode = 0
	response.Message = "berhasil login"
	response.Token = token

	return e.JSON(http.StatusOK, response)
}

func (c *UserDelivery) CreateUser(e echo.Context) error {

	ctx := e.Request().Context()
	form := new(users.CreateUserRequest)

	if err := e.Bind(form); err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	user, err := c.userUc.CreateUser(ctx, form)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err)
	}

	response := new(utils.Response)
	response.Message = "berhasil Create data"
	response.Data = user

	return e.JSON(http.StatusOK, response)
}
