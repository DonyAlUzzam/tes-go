package src

import (
	"fmt"
	"log"
	"myapp/config"
	v1 "myapp/src/delivery/v1"
	"myapp/src/usecase"
	"os"

	"github.com/labstack/echo/v4"
)

func Run(cfg config.Config) {

	echoServer := echo.New()

	us := usecase.NewUsecase(cfg)

	userDelivery := v1.NewUserDelivery(us.User())
	userGroup := echoServer.Group(os.Getenv("SERVICE_NAME") + `/v1/user`)
	userDelivery.Mount(userGroup)

	if err := echoServer.Start(fmt.Sprintf(":%s", os.Getenv("PORT_SERVER"))); err != nil {
		log.Panic(err)
	}
}
