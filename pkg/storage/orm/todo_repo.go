package orm

import (
	"errors"
	"fmt"

	"todo-lists/pkg/common"
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

func (t *todoRepo) Delete(id string, userId string) error {
	t.log.Debugf("deleting the todo with id : %s", id)

	todo := &todo.Todo{}
	err := t.db.Where("user_id = ? AND id = ?", userId, id).First(&todo).Error

	if err != nil {
		t.log.Debug("no single todo found")
		return err
	}

	if t.db.Delete(&todo, "id = ?", id).Error != nil {
		errMsg := fmt.Sprintf("error while deleting the user with id : %s", id)
		t.log.Errorf(errMsg)
		return errors.New(errMsg)
	}
	return nil
}

func (t *todoRepo) GetAll(userId string, filter common.FilterParam) ([]*todo.Todo, error) {
	t.log.Debug("get all the todos")

	todos := make([]*todo.Todo, 0)
	err := t.db.Limit(filter.Limit).Offset(filter.Offset).Find(&todos).Where("user_id = ?", userId).Error
	if err != nil {
		t.log.Debug("no single todo found")
		return nil, err
	}
	return todos, nil
}

func (t *todoRepo) GetByID(id string, userId string) (*todo.Todo, error) {
	t.log.Debugf("get todo details by id : %s", id)

	todo := &todo.Todo{}
	err := t.db.Where("user_id = ? AND id = ?", userId, id).First(&todo).Error
	if err != nil {
		t.log.Errorf("todo not found with id : %s, reason : %v", id, err)
		return nil, err
	}
	return todo, nil
}

func (t *todoRepo) Store(todo *todo.Todo) error {
	t.log.Debugf("creating a todo with user id : %v", todo.UserID)

	err := t.db.Create(&todo).Error
	if err != nil {
		t.log.Errorf("error while creating a todo, reason : %v", err)
		return err
	}
	return nil
}

func (t *todoRepo) Update(id string, tdo *todo.Todo) error {
	t.log.Debugf("updating the todo, user_id : %v", tdo.UserID)

	err := t.db.Model(&tdo).Where("user_id = ? AND id = ?", tdo.UserID, id).Update("content", tdo.Content).Error
	if err != nil {
		t.log.Errorf("error while updating a todo, reason : %v", err)
		return err
	}
	return nil
}