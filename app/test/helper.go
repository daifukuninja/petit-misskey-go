package test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/daifukuninja/petit-misskey-go/app/config"
	"github.com/stretchr/testify/require"
)

func NewConfig(t *testing.T, configType config.ConfigType) *config.Config {
	t.Helper()

	wd, _ := os.Getwd()
	cwd := filepath.Dir(wd)
	dirs := strings.Split(cwd, "/")
	traceBack := make([]string, 0)
	for i := len(dirs) - 1; i >= 0; i-- {
		traceBack = append(traceBack, "..")
		if dirs[i] == "app" {
			break
		}
	}

	require.Nil(t, os.Chdir(filepath.Join(traceBack...)))

	config, err := config.NewConfig(configType)

	require.NoError(t, err)

	return config
}
