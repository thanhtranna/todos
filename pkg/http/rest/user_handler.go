package rest

import (
	"todo-lists/pkg/common"
	"todo-lists/pkg/logger"
	"todo-lists/pkg/user"
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type userCtrl struct {
	log logger.LogInfoFormat
	svc user.Service
}

func NewUserCtrl(log logger.LogInfoFormat, svc user.Service) *userCtrl {
	return &userCtrl{log, svc}
}

func (u *userCtrl) GetAll(ctx *gin.Context) {
	users, err := u.svc.GetAll()
	if len(users) == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, common.ResponseWithError(err.Error(), nil))
		return
	}
	ctx.JSON(http.StatusOK, common.ResponseSuccess(users))
}

func (u *userCtrl) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	if _, err := uuid.FromString(id); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ResponseWithError(err.Error(), nil))
		return
	}

	user, err := u.svc.GetByID(id)
	if user == nil || err != nil {
		ctx.JSON(http.StatusNotFound, common.ResponseWithError(err.Error(), nil))
		return
	}
	ctx.JSON(http.StatusOK, common.ResponseSuccess(user))
}

func (u *userCtrl) Store(ctx *gin.Context) {
	var user user.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ResponseWithError(err.Error(), nil))
		return
	}
	err := u.svc.Store(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.ResponseWithError(err.Error(), nil))
		return
	}
	ctx.JSON(http.StatusCreated, common.ResponseSuccess(user))
}

func (u *userCtrl) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	if _, err := uuid.FromString(id); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ResponseWithError(err.Error(), nil))
		return
	}

	var user user.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ResponseWithError(err.Error(), nil))
		return
	}
	user.ID = id
	u.svc.Update(&user)
	ctx.JSON(http.StatusOK, common.ResponseSuccess(user))
}

func (u *userCtrl) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if _, err := uuid.FromString(id); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ResponseWithError(err.Error(), nil))
		return
	}
	u.svc.Delete(id)
	ctx.JSON(http.StatusOK, common.ResponseSuccess(nil))
}