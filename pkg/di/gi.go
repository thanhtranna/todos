package di

import (
	"todo-lists/pkg/config"
	"todo-lists/pkg/logger"
	"todo-lists/pkg/login"
	"todo-lists/pkg/storage"
	"todo-lists/pkg/storage/orm"
	"todo-lists/pkg/user"

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

	// login
	container.Provide(orm.NewLoginRepo)
	container.Provide(login.NewLoginService)

	// user
	container.Provide(orm.NewUserRepo)
	container.Provide(user.NewUserService)
	return container
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}