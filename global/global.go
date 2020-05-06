package global

import (
	"hello/model"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var (
	DB     *gorm.DB
	Conf   *viper.Viper
	Server model.Server
)
