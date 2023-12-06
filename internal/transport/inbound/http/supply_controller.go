package http_in

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"one-stop/internal/config"
	"one-stop/internal/service"
)

type userController struct {
	log *logrus.Logger
	env *config.Environment
	svc *service.Service
}

func (uc userController) CreateUserEntry(ctx *gin.Context) {
	fmt.Println("HELELLLOOOLOL")

}
