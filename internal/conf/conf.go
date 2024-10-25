package conf

import (
	"time"

	"github.com/spf13/viper"
)

type Conf struct {
	Server   *Server   `json:"server"`
	Data     *Data     `json:"data"`
	Security *Security `json:"security"`
}

type Security struct {
	JWT *JWT `json:"jwt"`
}

type JWT struct {
	Key string `json:"key"`
}

type Server struct {
	Addr    string        `json:"addr"`
	Timeout time.Duration `json:"timeout"`
}

type Data struct {
	Database *Database `json:"database,omitempty"`
	Redis    *Redis    `json:"redis,omitempty"`
}

type Database struct {
	Driver string `json:"driver"`
	Source string `json:"source"`
}

type Redis struct {
	Addr        string        `json:"addr,omitempty"`
	Password    string        `json:"password,omitempty"`
	DB          int           `json:"db"`
	ReadTimeout time.Duration `json:"read_timeout,omitempty"`
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
