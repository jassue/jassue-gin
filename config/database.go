package config

type Database struct {
    Driver string `mapstructure:"driver" json:"driver" yaml:"driver"`
    Host string `mapstructure:"host" json:"host" yaml:"host"`
    Port int `mapstructure:"port" json:"port" yaml:"port"`
    Database string `mapstructure:"database" json:"database" yaml:"database"`
    UserName string `mapstructure:"username" json:"username" yaml:"username"`
    Password string `mapstructure:"password" json:"password" yaml:"password"`
    Charset string `mapstructure:"charset" json:"charset" yaml:"charset"`
    MaxIdleConns int `mapstructure:"max_idle_conns" json:"max_idle_conns" yaml:"max_idle_conns"`
    MaxOpenConns int `mapstructure:"max_open_conns" json:"max_open_conns" yaml:"max_open_conns"`
    LogMode string `mapstructure:"log_mode" json:"log_mode" yaml:"log_mode"`
    EnableFileLogWriter bool `mapstructure:"enable_file_log_writer" json:"enable_file_log_writer" yaml:"enable_file_log_writer"`
    LogFilename string `mapstructure:"log_filename" json:"log_filename" yaml:"log_filename"`
}
