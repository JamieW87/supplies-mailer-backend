package errorhandling

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func HandleError(log *logrus.Logger, c *gin.Context, status int, msg string, err error) {

	log.Error(err)

	resp := ErrorResponse{Error: msg}
	c.IndentedJSON(status, resp)
}
