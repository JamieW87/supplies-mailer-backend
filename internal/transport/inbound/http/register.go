package http_in

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"one-stop/internal/config"
	"one-stop/internal/service"
)

func Register(r *gin.Engine, env *config.Environment, log *logrus.Logger, svc *service.Service) {

	publicRoutes := r.Group("/")
	uc := userController{env: env, log: log, svc: svc}

	publicRoutes.POST("/api/users/create", uc.CreateUserEntry)
	publicRoutes.GET("/health-check", HealthCheck)

}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{})
}
