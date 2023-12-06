package http_in

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"one-stop/internal/config"
)

type userController struct {
	log *logrus.Logger
	env *config.Environment
}

func (uc userController) CreateUserEntry(ctx *gin.Context) {

}
