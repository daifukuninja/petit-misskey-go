package misskey_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/daifukuninja/petit-misskey-go/infrastructure/misskey"
	"github.com/daifukuninja/petit-misskey-go/infrastructure/setting"
	model "github.com/daifukuninja/petit-misskey-go/model/misskey"
	"github.com/daifukuninja/petit-misskey-go/test"
	"github.com/daifukuninja/petit-misskey-go/util"
)

func TestMeta(t *testing.T) {
	config := test.NewConfig(t)
	setting := setting.NewUserSetting()
	instance := setting.GetInstanceByKey("io")

	client := misskey.NewClient(
		config,
		instance,
	)
	body := &model.Meta{
		AccessToken: model.AccessToken(config.Test.AccessToken),
		Detail:      false,
	}
	result, err := client.Meta(context.Background(), *body)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Println(util.PrittyJson(result))
}

func TestCreateNote(t *testing.T) {
	config := test.NewConfig(t)
	setting := setting.NewUserSetting()
	instance := setting.GetInstanceByKey("io")

	client := misskey.NewClient(
		config,
		instance,
	)
	unixnow := time.Now().Unix()
	text := `テスト投稿:desuwayo::aramaki:
	$[unixtime %d]
	`
	body := &model.CreateNote{
		AccessToken: model.AccessToken(config.Test.AccessToken),
		Visibility:  model.VisibilityHome,
		Text:        fmt.Sprintf(text, unixnow),
	}
	result, err := client.CreateNote(context.Background(), *body)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Println(util.PrittyJson(result))
}
