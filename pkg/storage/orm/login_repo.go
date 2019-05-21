package orm

import (
	"errors"

	"todo-lists/pkg/logger"
	"todo-lists/pkg/login"
	"todo-lists/pkg/user"
	"todo-lists/pkg/auth"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type loginRepo struct {
	db  *gorm.DB
	log logger.LogInfoFormat
	authSrv auth.AuthService
}

func NewLoginRepo(db *gorm.DB, log logger.LogInfoFormat, authSrv auth.AuthService) login.Repository {
	return &loginRepo{db, log, authSrv}
}

func (l *loginRepo) Signin(ln *login.Login) (string, error) {
	user := &user.User{}
	err := l.db.Where("email = ?", ln.Email).First(&user).Error
	if err != nil {
		l.log.Errorf("signin failed with reason : %v", err)
		return "", errors.New("something went wrong")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(ln.Password)); err != nil {
		return "", errors.New("password incorrect")
	}

	token, err := l.authSrv.GenerateToken(user.Email, user.ID)
	if err != nil {
		return "", errors.New("something went wrong")
	}

	return token, nil
}