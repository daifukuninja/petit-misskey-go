package router

import (
	"context"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/daifukuninja/petit-misskey-go/config"
	"github.com/daifukuninja/petit-misskey-go/infrastructure/setting"
	metaservice "github.com/daifukuninja/petit-misskey-go/service/meta"
	metamodel "github.com/daifukuninja/petit-misskey-go/view/meta"
)

type (
	// sub serviceを使ってtea programの実行制御を行うservice hub層
	Router struct {
		cfg         *config.Config
		userSetting *setting.UserSetting
	}
)

func New(cfg *config.Config, userSetting *setting.UserSetting) *Router {
	return &Router{
		cfg:         cfg,
		userSetting: userSetting,
	}
}

func (s *Router) Meta(ctx context.Context, instanceKey string) {
	instance := s.userSetting.GetInstanceByKey(instanceKey)
	if instance == nil {
		// TODO: ちゃんとエラー処理
		panic(fmt.Sprintf("instance key %s is not define", instanceKey))
	}
	service := metaservice.NewService(s.cfg, instance)

	model := metamodel.NewModel(ctx, service)

	if _, err := tea.NewProgram(model).Run(); err != nil {
		fmt.Println("meta error.") // TODO: ちゃんとエラー処理
		os.Exit(1)
	}

}
