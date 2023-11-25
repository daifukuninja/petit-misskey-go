package misskey_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/daifukuninja/petit-misskey-go/model/misskey"
	"github.com/daifukuninja/petit-misskey-go/test"
)

func TestMeta(t *testing.T) {
	config := test.NewConfig(t, "local")

	meta := &misskey.Meta{
		AccessToken: config.Misskey.AccessToken,
		Detail:      false,
	}
	j, err := json.Marshal(meta)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", j)
}
