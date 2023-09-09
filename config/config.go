package config

type Configuration struct {
	App App `mapstructure:"app" json:"app" yaml:"app"`
	Log Log `mapstructure:"Log" json:"Log" yaml:"Log"`
}
