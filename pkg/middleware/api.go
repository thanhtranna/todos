package middleware

import (
	"strings"
	"net/http"
	"fmt"

	"todo-lists/pkg/common"

	"github.com/gin-gonic/gin"
	jwt "github.com/dgrijalva/jwt-go"
)

type ApiMiddleware struct {}


func (m *ApiMiddleware) VerifyToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get token from header
		tokenString := strings.TrimSpace(ctx.GetHeader("Authorization"))
		tokenString = strings.Replace(tokenString, "Bearer ", "", -1)
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, common.ResponseWithError("User is not logged in", nil))
            ctx.Abort()
            return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:                                                                                                                                      
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return false, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			fmt.Println(token.Header)
			fmt.Println("dkmmm")
			return token.Header["kid"], nil
		})

		fmt.Println("err ", err)
		
		if err != nil {
			fmt.Println("asdasdasd")
			ctx.JSON(http.StatusUnauthorized, common.ResponseWithError(err.Error(), nil))
            ctx.Abort()
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid  {
			fmt.Println("claims", claims)
			userId := claims["sub"].(string)
			fmt.Println("userId ", userId)
		}
	
		
		// Verify token
		// authSrv := auth.NewAuthService()

		// fmt.Println("header ", token)
		// ok, data := m.authSrv.ParseToken(token)

		// fmt.Println("ok ", ok)
		// fmt.Println("data ", data)
		ctx.Next()
	}
}

