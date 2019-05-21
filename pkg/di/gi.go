package di

import (
	"todo-lists/pkg/config"
	"todo-lists/pkg/logger"
	"todo-lists/pkg/auth"
	"todo-lists/pkg/login"
	"todo-lists/pkg/storage"
	"todo-lists/pkg/storage/orm"
	"todo-lists/pkg/user"
	"todo-lists/pkg/todo"

	"go.uber.org/dig"
)

var container = dig.New()

func BuildContainer() *dig.Container {
	// config
	container.Provide(config.NewConfig)

	// DB
	container.Provide(storage.NewDb)

	// logger
	container.Provide(logger.NewLogger)

	// auth
	container.Provide(auth.NewAuthService)

	// login
	container.Provide(orm.NewLoginRepo)
	container.Provide(login.NewLoginService)

	// user
	container.Provide(orm.NewUserRepo)
	container.Provide(user.NewUserService)

	// todo
	container.Provide(orm.NewTodoRepo)
	container.Provide(todo.NewTodoService)
	return container
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}