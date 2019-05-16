package orm

import (
	"errors"
	"fmt"

	"todo-lists/pkg/logger"
	"todo-lists/pkg/todo"

	"github.com/jinzhu/gorm"
)

type todoRepo struct {
	db  *gorm.DB
	log logger.LogInfoFormat
}

func NewTodoRepo(db *gorm.DB, log logger.LogInfoFormat) todo.Repository {
	return &todoRepo{db, log}
}

func (u *todoRepo) Delete(id string, userId string) error {
	u.log.Debugf("deleting the todo with id : %s", id)

	if u.db.Delete(&todo.Todo{}, "id = ? AND user_id = ?", id, userId).Error != nil {
		errMsg := fmt.Sprintf("error while deleting a todo with id : %s", id)
		u.log.Errorf(errMsg)
		return errors.New(errMsg)
	}
	return nil
}

func (u *todoRepo) GetAll() ([]*todo.Todo, error) {
	u.log.Debug("get all the todos")

	todos := make([]*todo.Todo, 0)
	err := u.db.Find(&todos).Error
	if err != nil {
		u.log.Debug("no single todo found")
		return nil, err
	}
	return users, nil
}

func (u *todoRepo) GetByID(id string) (*user.User, error) {
	u.log.Debugf("get user details by id : %s", id)

	user := &user.User{}
	err := u.db.Where("user_id = ?", id).First(&user).Error
	if err != nil {
		u.log.Errorf("user not found with id : %s, reason : %v", id, err)
		return nil, err
	}
	return user, nil
}

func (u *todoRepo) Store(usr *user.User) error {
	u.log.Debugf("creating the user with email : %v", usr.Email)

	err := u.db.Create(&usr).Error
	if err != nil {
		u.log.Errorf("error while creating the user, reason : %v", err)
		return err
	}
	return nil
}

func (u *todoRepo) Update(usr *user.User) error {
	u.log.Debugf("updating the user, user_id : %v", usr.ID)

	err := u.db.Model(&usr).Updates(user.User{FirstName: usr.FirstName, LastName: usr.LastName, Password: usr.Password}).Error
	if err != nil {
		u.log.Errorf("error while updating the user, reason : %v", err)
		return err
	}
	return nil
}