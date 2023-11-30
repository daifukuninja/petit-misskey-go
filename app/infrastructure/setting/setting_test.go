package setting_test

import (
	"testing"

	"github.com/daifukuninja/petit-misskey-go/infrastructure/setting"
	"github.com/stretchr/testify/assert"
)

func TestWriteValue(t *testing.T) {
	i := &setting.Instance{
		BaseUrl:     "https://misskey.io/api",
		UserName:    "wasya",
		AccessToken: "ZzkCk14l33aOCndxvHMwbsRhXMpFdAO5",
	}
	imap := make(map[string]setting.Instance)
	imap["io"] = *i

	setting := setting.NewUserSetting()

	setting.WriteValue(imap)

	instances := setting.GetInstances()

	assert.True(t, len(instances) > 0)

	instance := setting.GetInstanceByKey("io")

	assert.NotNil(t, instance)
	assert.Equal(t, i.BaseUrl, instance.BaseUrl)
	assert.Equal(t, i.UserName, instance.UserName)
	assert.Equal(t, i.AccessToken, instance.AccessToken)
}
