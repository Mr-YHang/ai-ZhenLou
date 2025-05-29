package global

import (
	"ai-ZhenLou/config"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type Application struct {
	Config *config.Config // 全局配置文件
	Log    zerolog.Logger // 全局日志
	DB     *gorm.DB       // 全局mysql
}

var App = new(Application)
