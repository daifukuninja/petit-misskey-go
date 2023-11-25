package misskey_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/daifukuninja/petit-misskey-go/infrastructure/misskey"
	model "github.com/daifukuninja/petit-misskey-go/model/misskey"
	"github.com/daifukuninja/petit-misskey-go/test"
)

func TestMeta(t *testing.T) {
	config := test.NewConfig(t, "local")

	client := misskey.NewClient(
		config.Misskey.BaseUrl,
		config.Http.Timeout,
	)
	body := &model.Meta{
		AccessToken: config.Misskey.AccessToken,
		Detail:      false,
	}
	result, err := client.Meta(context.Background(), *body)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("%v\n", result)
}

func TestCreateNote(t *testing.T) {
	config := test.NewConfig(t, "local")
	client := misskey.NewClient(
		config.Misskey.BaseUrl,
		config.Http.Timeout,
	)
	body := &model.CreateNote{
		AccessToken: config.Misskey.AccessToken,
		Visibility:  model.VisibilityPublic,
		Text:        "Goでテスト投稿 :aramaki:",
	}
	result, err := client.CreateNote(context.Background(), *body)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("%v\n", result)
}
