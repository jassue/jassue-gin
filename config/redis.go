package config

type Redis struct {
    Host string `mapstructure:"host" json:"host" yaml:"host"`
    Port string `mapstructure:"port" json:"port" yaml:"port"`
    DB int `mapstructure:"db" json:"db" yaml:"db"`
    Password string `mapstructure:"password" json:"password" yaml:"password"`
}
