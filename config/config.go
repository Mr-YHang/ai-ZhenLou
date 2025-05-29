package config

var Conf = new(Config)

type Config struct {
	App App `mapstructure:"app" json:"app" yaml:"app"`
	Log Log `mapstructure:"log" json:"log" yaml:"log"`
}

type App struct {
	Env     string `mapstructure:"env" json:"env" yaml:"env"`                // 环境
	Port    string `mapstructure:"port" json:"port" yaml:"port"`             // 端口
	AppName string `mapstructure:"app_name" json:"app_name" yaml:"app_name"` // 应用名
	AppUrl  string `mapstructure:"app_url" json:"app_url" yaml:"app_url"`    // 域名
}

type Log struct {
	IsConsole  bool   `mapstructure:"is_console" json:"is_console" yaml:"is_console"`    // 命令台显示
	Path       string `mapstructure:"path" json:"path" yaml:"path"`                      // 日志路径
	MaxAge     int    `mapstructure:"max_age" json:"max_age" yaml:"max_age"`             // 保留天数
	MaxSize    int    `mapstructure:"max_size" json:"max_size" yaml:"max_size"`          // 单文件最大M
	Level      int    `mapstructure:"level" json:"level" yaml:"level"`                   // 等级
	MaxBackups int    `mapstructure:"max_backups" json:"max_backups" yaml:"max_backups"` // 最大文件数
}
