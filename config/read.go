package config

import (
	"fmt"
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"github.com/zsakvo/Cirno-go/structure"
)

var Config structure.ConfigStruct

func InitConfig(hasConfig bool) {
	if hasConfig {
		dir, _ := homedir.Dir()
		expandedDir, _ := homedir.Expand(dir)
		viper.SetConfigName("config")
		viper.AddConfigPath(expandedDir + "/Cirno")
		viper.SetConfigType("yaml")
		viper.SetDefault("app.app_version", "2.1.3")
		viper.SetDefault("app.device_token", "duyuedu_")
		viper.SetDefault("app.user_agent", "Duyuedu/2.1.3 (iPad; iOS 15.6; Scale/2.00)")
		viper.SetDefault("app.default_key", "Gnzg4OfHHp103jx8ebG82gA8VPJZnLTk")
		viper.SetDefault("app.host_url", "https://app.duread8.com")
		viper.SetDefault("extra.coroutines", 3)
		viper.SetDefault("extra.cpic", false)
		viper.SetDefault("extra.cache_no_paid", false)
		err := viper.ReadInConfig()
		if err != nil {
			fmt.Printf("config file error: %s\n", err)
			os.Exit(1)
		}
		err = viper.Unmarshal(&Config) // 将配置解析到 config 变量
		if err != nil {
			log.Fatalf("unable to decode into struct, %v", err)
		}
	} else {
		Config = structure.ConfigStruct{
			App: structure.App{
				AppVersion:  "2.1.3",
				DeviceToken: "duyuedu_",
				UserAgent:   "Duyuedu/2.1.3 (iPad; iOS 15.6; Scale/2.00)",
				DefaultKey:  "Gnzg4OfHHp103jx8ebG82gA8VPJZnLTk",
				HostUrl:     "http://app.duread8.com",
			},
		}
	}
}
