package meta

import (
	"context"

	"github.com/daifukuninja/petit-misskey-go/model/misskey"
)

type (
	Client interface {
		Meta(ctx context.Context, contents misskey.Meta) (*misskey.MetaResponse, error)
	}
)
