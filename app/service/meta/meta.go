package meta

import (
	"context"

	"github.com/daifukuninja/petit-misskey-go/config"
	"github.com/daifukuninja/petit-misskey-go/infrastructure/misskey"
	"github.com/daifukuninja/petit-misskey-go/infrastructure/setting"
	model "github.com/daifukuninja/petit-misskey-go/model/misskey"
)

type Service struct {
	client      *misskey.Client
	accessToken model.AccessToken
}

func NewService(cfg config.Config, instance setting.Instance) *Service {
	client := misskey.NewClient(instance.BaseUrl, cfg.Http.Timeout)

	return &Service{
		client: client,
	}
}

func (s *Service) Do(ctx context.Context) (*model.MetaResponse, error) {
	contents := &model.Meta{
		AccessToken: s.accessToken,
		Detail:      false,
	}
	res, clientErr := s.client.Meta(ctx, *contents)
	if clientErr != nil {
		return nil, clientErr
	}
	return res, nil
}
