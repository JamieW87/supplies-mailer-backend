package http_in

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"one-stop/internal/config"
	errorhandling "one-stop/internal/errors"
	"one-stop/internal/model"
	"one-stop/internal/service"
)

type userController struct {
	log *logrus.Logger
	env *config.Environment
	svc *service.Service
}

func (uc userController) CreateUserEntry(c *gin.Context) {

	var req model.CreateUserEntryRequest
	err := c.BindJSON(&req)
	if err != nil {
		errorhandling.HandleError(uc.log, c, http.StatusInternalServerError, "Oops, something went wrong", err)
		return
	}

	u, err := uc.svc.StoreUserData(c, req.Name, req.Email, req.Phone)
	if err != nil {
		errorhandling.HandleError(uc.log, c, http.StatusInternalServerError, "Oops, something went wrong", err)
		return
	}

	err = uc.svc.InsertUserCategory(c, u, req.Category)
	if err != nil {
		errorhandling.HandleError(uc.log, c, http.StatusInternalServerError, "Oops, something went wrong", err)
		return
	}

	fmt.Println("ayyeeeee")

}
