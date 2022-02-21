package config

import (
	"bytes"
	"github.com/spf13/viper"
	"strings"
)

func Read(config []byte) {
	viper.SetConfigType("yml")
	err := viper.ReadConfig(bytes.NewBuffer(config))
	if err != nil {
		panic(err)
	}
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	viper.AutomaticEnv()
}
