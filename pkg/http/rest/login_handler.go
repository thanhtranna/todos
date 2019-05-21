package rest

import (
	"net/http"

	"todo-lists/pkg/common"
	"todo-lists/pkg/logger"
	"todo-lists/pkg/login"

	"github.com/gin-gonic/gin"
)

type loginCtrl struct {
	log logger.LogInfoFormat
	svc login.Service
}

func NewLoginCtrl(log logger.LogInfoFormat, svc login.Service) *loginCtrl {
	return &loginCtrl{log, svc}
}

func (l loginCtrl) Signin(ctx *gin.Context) {
	var lg login.Login
	if err := ctx.ShouldBindJSON(&lg); err != nil {
		l.log.Errorf("signin request bind failed with reason : %v", err)
		ctx.JSON(http.StatusBadRequest, common.ResponseWithError(err.Error(), nil))
		return
	}

	token, err := l.svc.Signin(lg.Email, lg.Password)
	if err != nil {
		ctx.JSON(http.StatusForbidden, common.ResponseFail(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, common.ResponseSuccess(token))
}