//+build wireinject

package main

import (
	"geektime/internal/conf"
	"geektime/internal/data"
	"geektime/internal/server"
	"geektime/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
)

func initAPP(config *conf.MydemoConfig) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, service.ProviderSet, newApp))
}
