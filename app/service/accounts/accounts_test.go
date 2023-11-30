package accounts_test

import (
	"testing"

	"github.com/daifukuninja/petit-misskey-go/infrastructure/setting"
	"github.com/daifukuninja/petit-misskey-go/service/accounts"
	"github.com/stretchr/testify/assert"
)

func TestGetIO(t *testing.T) {
	setting := setting.NewUserSetting()
	accounts := accounts.NewService(setting)

	instance := accounts.Get("io")

	assert.NotNil(t, instance)
}

// TODO: 書き込みのテスト
