package resolver_test

import (
	"testing"

	"github.com/daifukuninja/petit-misskey-go/infrastructure/resolver"
	"github.com/stretchr/testify/assert"
)

func TestMisskeyResolver(t *testing.T) {
	resolver := resolver.NewMisskeyStreamUrlResolver()

	url, err := resolver.Resolve(
		"https://misskey.io/api",
		map[string]string{
			"accessToken": "test",
		})
	assert.NoError(t, err)
	assert.Equal(t, "wss://misskey.io/streaming?i=test", url)
}
