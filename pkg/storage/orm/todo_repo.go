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

func (u *todoRepo) GetAll(userId string) ([]*todo.Todo, error) {
	u.log.Debug("get all the todos")

	todos := make([]*todo.Todo, 0)
	err := u.db.Find(&todos).Where("user_id = ?", userId).Error
	if err != nil {
		u.log.Debug("no single todo found")
		return nil, err
	}
	return todos, nil
}

func (u *todoRepo) GetByID(id string, userId string) (*todo.Todo, error) {
	u.log.Debugf("get todo details by id : %s", id)

	todo := &todo.Todo{}
	err := u.db.Where("user_id = ? AND id = ?", userId, id).First(&todo).Error
	if err != nil {
		u.log.Errorf("todo not found with id : %s, reason : %v", id, err)
		return nil, err
	}
	return todo, nil
}

func (u *todoRepo) Store(todo *todo.Todo) error {
	u.log.Debugf("creating a todo with user id : %v", todo.UserID)

	err := u.db.Create(&todo).Error
	if err != nil {
		u.log.Errorf("error while creating a todo, reason : %v", err)
		return err
	}
	return nil
}

func (u *todoRepo) Update(tdo *todo.Todo) error {
	u.log.Debugf("updating the todo, user_id : %v", tdo.UserID)

	err := u.db.Model(&tdo).Updates(todo.Todo{Content: tdo.Content}).Error
	if err != nil {
		u.log.Errorf("error while updating a todo, reason : %v", err)
		return err
	}
	return nil
}