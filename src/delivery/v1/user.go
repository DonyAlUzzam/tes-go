package v1

import (
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
	group.GET("/detail", c.User)
}

func (c *UserDelivery) User(e echo.Context) error {

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
