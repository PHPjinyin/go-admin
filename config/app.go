package config

type App struct {
	Env       string `mapstructure:"app" json:"app" yaml:"app"`
	Port      string `mapstructure:"port" json:"port" yaml:"port"`
	AppName   string `mapstructure:"app_name" json:"app_name" yaml:"app_name"`
	AppUrl    string `mapstructure:"app_url" json:"app_url" yaml:"app_url"`
	StopDelay int    `mapstructure:"stop_delay",json:"stop_delay" yaml:"stop_delay"`
}
