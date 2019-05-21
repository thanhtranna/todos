package middleware

import (
	"strings"
	"net/http"
	"fmt"

	"todo-lists/pkg/common"
	"todo-lists/pkg/config"

	"github.com/gin-gonic/gin"
	jwt "github.com/dgrijalva/jwt-go"
)

type ApiMiddleware struct {}


func (m *ApiMiddleware) VerifyToken(config *config.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get token from header
		tokenString := strings.TrimSpace(ctx.GetHeader("Authorization"))
		tokenString = strings.Replace(tokenString, "Bearer ", "", -1)
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, common.ResponseWithError("User is not logged in", nil))
            ctx.Abort()
            return
		}

		// Verify token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:                                                                                                                                      
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return false, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(config.JWT.Key), nil
		})
		
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, common.ResponseWithError(err.Error(), nil))
            ctx.Abort()
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userInfo := common.UserInfo{
				UserID:		claims["userId"].(string),
				UserName:	claims["username"].(string),
			}
			ctx.Set("userInfo", userInfo)
			ctx.Next()
			return
		}
		ctx.JSON(http.StatusUnauthorized, common.ResponseWithError(err.Error(), nil))
		ctx.Abort()
		return
	}
}

