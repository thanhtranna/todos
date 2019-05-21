package todo

import "todo-lists/pkg/common"

type Repository interface {
	Delete(id string, userId string) error
	GetAll(userId string, filter common.FilterParam) ([]*Todo, error)
	GetByID(id string, userId string) (*Todo, error)
	Store(t *Todo) error
	Update(id string, t *Todo) error
}