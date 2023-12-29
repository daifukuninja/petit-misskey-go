package websocket_test

import (
	"os"
	"testing"

	"github.com/daifukuninja/petit-misskey-go/infrastructure/resolver"
	"github.com/daifukuninja/petit-misskey-go/infrastructure/websocket"
	"github.com/daifukuninja/petit-misskey-go/test"
)

func TestGetStream(t *testing.T) {
	cfg := test.NewConfig(t)
	resolver := resolver.NewMisskeyStreamUrlResolver()
	wsClient := websocket.NewClient(cfg.Test.BaseUrl, cfg.Test.AccessToken, resolver, os.Stdout)
	wsClient.Start()
}
