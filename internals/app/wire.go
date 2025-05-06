//go:build wireinject
// +build wireinject

package app

// The build tag makes sure the stub is not built in the final build.

import (
	"FliQt/internals/app/api"
	"FliQt/internals/app/config"
	"FliQt/internals/app/repository"
	"FliQt/internals/app/router"
	"FliQt/internals/app/service"
	"FliQt/pkg/di"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

type Application struct {
	Engine *gin.Engine
}

func injector(cfg *config.Config) (*Application, func(), error) {
	wire.Build(
		repository.WireSet,
		service.WireSet,
		api.WireSet,
		router.New,
		di.WireSet,
		wire.Struct(new(Application), "*"),
	)
	return nil, nil, nil
}

func Start(cfg *config.Config) (*Application, func(), error) {
	return injector(cfg)
}
