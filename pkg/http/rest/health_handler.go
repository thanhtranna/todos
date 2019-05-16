package rest

import (
	"net/http"

	"todo-lists/pkg/common"
	
	"github.com/gin-gonic/gin"
)

type healthCtrl struct{}

func NewHealthCtrl() *healthCtrl {
	return &healthCtrl{}
}

func (h healthCtrl) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, common.ResponseSuccess("health check"))
}