package conf

import (
	"time"

	"github.com/spf13/viper"
)

type Conf struct {
	Env      string    `mapstructure:"env"`
	Server   *Server   `mapstructure:"server"`
	Data     *Data     `mapstructure:"data"`
	Security *Security `mapstructure:"security"`
	Log      *Log      `mapstructure:"log"`
}

type Log struct {
	LogLevel    string `mapstructure:"log_level"`
	MaxSize     int    `mapstructure:"max_size"`
	MaxAge      int    `mapstructure:"max_age"`
	Compress    bool   `mapstructure:"compress"`
	LogFileName string `mapstructure:"log_file_name"`
	MaxBackup   int    `mapstructure:"max_backup"`
	Encoding    string `mapstructure:"encoding"`
}

type Security struct {
	JWT *JWT `json:"jwt"`
}

type JWT struct {
	Key string `json:"key"`
}

type Server struct {
	Addr    string        `mapstructure:"addr"`
	Timeout time.Duration `mapstructure:"timeout"`
}

type Data struct {
	Database *Database `mapstructure:"database,omitempty"`
	Redis    *Redis    `mapstructure:"redis,omitempty"`
}

type Database struct {
	Driver string `mapstructure:"driver"`
	Source string `mapstructure:"source"`
}

type Redis struct {
	Addr        string        `mapstructure:"addr,omitempty"`
	Password    string        `mapstructure:"password,omitempty"`
	DB          int           `mapstructure:"db"`
	ReadTimeout time.Duration `mapstructure:"read_timeout,omitempty"`
}

func New() (*Conf, error) {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./configs")

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Conf
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
