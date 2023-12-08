//go:build wireinject
// +build wireinject

package di

import (
	"github.com/daifukuninja/petit-misskey-go/config"
	"github.com/daifukuninja/petit-misskey-go/infrastructure/setting"
	"github.com/daifukuninja/petit-misskey-go/router"
	"github.com/google/wire"
)

func InitializeRouter() *router.Router {
	wire.Build(
		router.New,
		config.ProviderSet,
		setting.ProviderSet,
	)
	return &router.Router{}
}
