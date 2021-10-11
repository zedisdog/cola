package cola

import (
	"bytes"
	"github.com/spf13/viper"
	"os"
	"strings"
)

// Config 配置
//  config: 配置文件内容的bytes,yml格式
func Config(config []byte) {
	// 读config
	viper.SetConfigType("yml")
	err := viper.ReadConfig(bytes.NewBuffer(config))
	if err != nil {
		panic(err)
	}

	// 环境变量中的__换成. 然后读环境变量，覆盖到配置
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	viper.AutomaticEnv()

	// 设置内置配置
	// 设置根
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	viper.Set("root", wd)
}
