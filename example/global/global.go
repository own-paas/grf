package global

import (
	"github.com/sestack/grf/example/config"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	VP     *viper.Viper
	CONFIG config.GlobalConfig
)
