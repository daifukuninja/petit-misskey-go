//go:build wireinject
// +build wireinject

package meta

import (
	"github.com/daifukuninja/petit-misskey-go/config"
	"github.com/daifukuninja/petit-misskey-go/domain/meta"
	"github.com/daifukuninja/petit-misskey-go/infrastructure/bubbles"
	"github.com/daifukuninja/petit-misskey-go/infrastructure/misskey"
	"github.com/daifukuninja/petit-misskey-go/infrastructure/setting"
	service "github.com/daifukuninja/petit-misskey-go/service/meta"
	"github.com/google/wire"
)

func InitializeModel(instance *setting.Instance) *Model {
	wire.Build(
		NewModel,
		service.NewService,
		config.NewConfig,
		misskey.NewClient,
		bubbles.ProviderSet,
		wire.Bind(new(meta.Client), new(*misskey.Client)),
		wire.Bind(new(bubbles.SimpleViewFactory), new(*bubbles.ViewportFactory)),
	)
	return &Model{}
}
