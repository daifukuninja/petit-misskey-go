package config

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/daifukuninja/petit-misskey-go/app/model/misskey"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var (
	instance = new(Config)
	once     sync.Once
)

// /config/config.{type}.yaml のファイルを読み込む
type ConfigType string

type Config struct {
	Misskey struct {
		AccessToken misskey.AccessToken
		BaseUrl     string
	}
	Http struct {
		Timeout time.Duration
	}
}

func NewConfig(configType ConfigType) (*Config, error) {
	var err error
	once.Do(func() {
		err = readConfig(configType)
	})
	if err != nil {
		return nil, err
	}
	return instance, nil
}

func readConfig(configType ConfigType) error {
	cwd, _ := os.Getwd()
	fmt.Printf("%s\n", cwd)

	name := fmt.Sprintf("config.%s", configType)
	viper.SetConfigName(name)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config/")

	if err := viper.ReadInConfig(); err != nil {
		return errors.WithStack(err)
	}
	if err := viper.Unmarshal(&instance); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
