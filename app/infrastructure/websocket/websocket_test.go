package websocket_test

import (
	"testing"

	"github.com/daifukuninja/petit-misskey-go/infrastructure/websocket"
	"github.com/daifukuninja/petit-misskey-go/test"
)

func TestGetStream(t *testing.T) {
	cfg := test.NewConfig(t)
	wsClient := websocket.NewClient(cfg.Test.BaseUrl, cfg.Test.AccessToken)
	wsClient.Start()
}
