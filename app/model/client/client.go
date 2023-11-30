package client

import (
	"context"

	"github.com/daifukuninja/petit-misskey-go/model/misskey"
)

type (
	Meta interface {
		Meta(ctx context.Context, contents misskey.Meta) (*misskey.MetaResponse, error)
	}
)
