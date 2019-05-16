package middleware

import (
	"strings"
	"net/http"
	"todo-lists/pkg/common"

	"github.com/gin-gonic/gin"
)

type Api struct {}


func (m *Api) VerifyToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get token from header
		token := strings.TrimSpace(c.GetHeader("Authorization"))
		if token == "" {
			ctx.JSON(http.StatusBadRequest, common.ResponseWithError("User is not logged in", nil))
            ctx.Abort()
            return
		}
		
		// Verify token
	}
}