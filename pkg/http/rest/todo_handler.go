package rest

import (
	"net/http"

	"todo-lists/pkg/common"
	"todo-lists/pkg/logger"
	"todo-lists/pkg/todo"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type todoCtrl struct {
	log logger.LogInfoFormat
	svc todo.Service
}

func NewTodoCtrl(log logger.LogInfoFormat, svc todo.Service) *todoCtrl {
	return &todoCtrl{log, svc}
}

func (t *todoCtrl) GetAll(ctx *gin.Context) {
	userInfoData, ok := ctx.Get("userInfo")
	if !ok {
		ctx.JSON(http.StatusBadRequest, common.ResponseWithError("user not accept", nil))
		return
	}

	page, limit := common.GetPagination(ctx)
	skip := common.GetSkipValue(page, limit)
	filter := common.FilterParam{
		Limit: limit,
		Page: page,
		Offset:	skip,
	}

	userInfo := userInfoData.(common.UserInfo)

	todos, err := t.svc.GetAll(userInfo.UserID, filter)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.ResponseWithError(err.Error(), nil))
		return
	}
	ctx.JSON(http.StatusOK, common.ResponseSuccess(todos))
}

func (t *todoCtrl) GetByID(ctx *gin.Context) {
	userInfoData, ok := ctx.Get("userInfo")
	if !ok {
		ctx.JSON(http.StatusBadRequest, common.ResponseWithError("user not accept", nil))
		return
	}
	id := ctx.Param("id")
	if _, err := uuid.FromString(id); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ResponseWithError(err.Error(), nil))
		return
	}

	userInfo := userInfoData.(common.UserInfo)

	user, err := t.svc.GetByID(id, userInfo.UserID)
	if user == nil || err != nil {
		ctx.JSON(http.StatusNotFound, common.ResponseWithError(err.Error(), nil))
		return
	}
	ctx.JSON(http.StatusOK, common.ResponseSuccess(user))
}

func (t *todoCtrl) Store(ctx *gin.Context) {
	userInfoData, ok := ctx.Get("userInfo")
	if !ok {
		ctx.JSON(http.StatusBadRequest, common.ResponseWithError("user not accept", nil))
		return
	}

	userInfo := userInfoData.(common.UserInfo)
	var todo todo.Todo
	if err := ctx.ShouldBindJSON(&todo); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ResponseWithError(err.Error(), nil))
		return
	}
	todo.UserID = userInfo.UserID
	err := t.svc.Store(&todo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.ResponseWithError(err.Error(), nil))
		return
	}
	ctx.JSON(http.StatusCreated, common.ResponseSuccess(todo))
}

func (t *todoCtrl) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	if _, err := uuid.FromString(id); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ResponseWithError(err.Error(), nil))
		return
	}

	var todo todo.Todo
	if err := ctx.ShouldBindJSON(&todo); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ResponseWithError(err.Error(), nil))
		return
	}
	if todo.ID != id {
		ctx.JSON(http.StatusBadRequest, common.ResponseWithError("not owner", nil))
		return
	}

	userInfoData, ok := ctx.Get("userInfo")
	if !ok {
		ctx.JSON(http.StatusBadRequest, common.ResponseWithError("user not accept", nil))
		return
	}

	userInfo := userInfoData.(common.UserInfo)
	if todo.UserID != userInfo.UserID {
		ctx.JSON(http.StatusBadRequest, common.ResponseWithError("not owner", nil))
		return
	}
	todo.UserID = userInfo.UserID
	err := t.svc.Update(id, &todo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.ResponseWithError(err.Error(), nil))
		return
	}
	ctx.JSON(http.StatusOK, common.ResponseSuccess(todo))
}

func (t *todoCtrl) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if _, err := uuid.FromString(id); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ResponseWithError(err.Error(), nil))
		return
	}
	userInfoData, ok := ctx.Get("userInfo")
	if !ok {
		ctx.JSON(http.StatusBadRequest, common.ResponseWithError("user not accept", nil))
		return
	}

	userInfo := userInfoData.(common.UserInfo)
	err := t.svc.Delete(id, userInfo.UserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.ResponseWithError(err.Error(), nil))
		return
	}
	ctx.JSON(http.StatusOK, common.ResponseSuccess(nil))
}