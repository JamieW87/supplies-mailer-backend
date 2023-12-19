package http_in

import (
	"errors"
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

	uc.log.Info("Create user request received")
	var req model.CreateUserEntryRequest
	err := c.BindJSON(&req)
	if err != nil {
		errorhandling.HandleError(uc.log, c, http.StatusInternalServerError, "Oops, something went wrong", err)
		return
	}
	
	uc.log.Info(fmt.Sprintf("Received request for user %s for category %s", req.Name, req.Category))

	if !uc.svc.IsValidCategory(req.Category) {
		errorhandling.HandleError(uc.log, c, http.StatusBadRequest, "Invalid category", errors.New("category does not exist"))
		return
	}

	uc.log.Info("Storing user data")
	u, err := uc.svc.StoreUserData(c, req.Name, req.Email, req.Phone, req.Category)
	if err != nil {
		errorhandling.HandleError(uc.log, c, http.StatusInternalServerError, "Oops, something went wrong", err)
		return
	}

	uc.log.Info("Getting supplier information")
	emails, err := uc.svc.RetrieveSupplierInfo(c, req.Category)
	if err != nil {
		errorhandling.HandleError(uc.log, c, http.StatusInternalServerError, "Oops, something went wrong", err)
		return
	}

	uc.log.Info("Sending email to suppliers")
	err = uc.svc.SendSupplierEmail(c, emails, req.Name, req.Email, req.Category)
	if err != nil {
		errorhandling.HandleError(uc.log, c, http.StatusInternalServerError, "Oops, something went wrong", err)
		return
	}

	c.IndentedJSON(http.StatusOK, &model.CreateUserEntryResponse{UserId: u.String()})
}
