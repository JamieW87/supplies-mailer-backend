package http_in

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"one-stop/internal/config"
)

func Register(r *gin.Engine, env *config.Environment, log *logrus.Logger) {

	publicRoutes := r.Group("/")

	uc := userController{env: env, log: log}

	publicRoutes.POST("/api/users/create", uc.CreateUserEntry)

}
