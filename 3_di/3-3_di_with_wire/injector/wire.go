//go:build wireinject

package injector

import (
    "github.com/google/wire"
    "myproject/controller"
    "myproject/infra"
    "myproject/repository"
    "myproject/usecase"
)

// db
var dbSet = wire.NewSet(
    infra.InitDB,
)

// repository
var repositorySet = wire.NewSet(
    repository.NewUserRepository,
)

// usecase
var usecaseSet = wire.NewSet(
    usecase.NewUserUsecase,
)

// controller
var controllerSet = wire.NewSet(
    controller.NewUserController,
)

func InitializeController() controller.UserController {
    wire.Build(
        dbSet,
        repositorySet,
        usecaseSet,
        controllerSet,
    )
    return nil
}

