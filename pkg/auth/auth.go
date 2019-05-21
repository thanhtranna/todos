package auth

import (
	"time"
	"fmt"

	"todo-lists/pkg/config"

	jwt "github.com/dgrijalva/jwt-go"
)

type AuthService interface {
	GenerateToken(email, userId string) (string, error)
	ParseToken(unparsedToken string) (bool, string)
}

type authService struct {
	config *config.Config
}


func NewAuthService(c *config.Config) AuthService {
	return &authService{
		config: c,
	}
}

func (avc *authService) GenerateToken(email, userId string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	// Create the Claims
	claims := token.Claims.(jwt.MapClaims)

	// set some claims                                                                                                                                                                                   
	claims["username"] = email;
	claims["userId"] = userId;
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	return token.SignedString([]byte(avc.config.JWT.Key))
}

func (avc *authService) ParseToken(unparsedToken string) (bool, string) {
	token, err := jwt.Parse(unparsedToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:                                                                                                                                      
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return false, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return token.Header["kid"], nil
	})

	if err == nil && token.Valid {
		return true, unparsedToken
	} else {
		return false, ""
	}
}