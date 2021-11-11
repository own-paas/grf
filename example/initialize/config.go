package initialize

import (
	"fmt"
	"github.com/sestack/grf/example/global"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/toolkits/file"
)

func InitConfig(config string) (*viper.Viper, error) {
	if config == "" {
		return nil, fmt.Errorf("use -c to specify configuration file")
	}

	if !file.IsExist(config) {
		return nil, fmt.Errorf("config file:%s is not existent.", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetDefault("db.driver", "mysql")
	v.SetDefault("db.dsn", "root:@(127.0.0.1)/example?charset=utf8&parseTime=true&loc=Local")
	v.SetDefault("http.ssl", false)
	v.SetDefault("http.address", "127.0.0.1:8888")

	err := v.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("Fatal error config file: %s \n", err)
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.CONFIG); err != nil {
		return nil, err
	}

	return v, nil
}
