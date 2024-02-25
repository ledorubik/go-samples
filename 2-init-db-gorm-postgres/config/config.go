package config

type Config struct {
	// DB config
	DbUserName             string `mapstructure:"db_username"`
	DbPassword             string `mapstructure:"db_password"`
	DbHost                 string `mapstructure:"db_host"`
	DbPort                 int    `mapstructure:"db_port"`
	DbName                 string `mapstructure:"db_name"`
	DbSchema               string `mapstructure:"db_schema"`
	DbPreferSimpleProtocol bool   `mapstructure:"db_prefer_simple_protocol"`
	DbPrepareStatement     bool   `mapstructure:"db_prepare_statement"`
	DbMigrate              bool   `mapstructure:"db_migrate"`
	DbLogLevel             string `mapstructure:"db_log_level"`
}

func New() *Config {
	return &Config{
		DbUserName:             "postgres",
		DbPassword:             "postgres",
		DbHost:                 "localhost",
		DbPort:                 5432,
		DbName:                 "go-samples",
		DbSchema:               "go-samples",
		DbPreferSimpleProtocol: false,
		DbPrepareStatement:     true,
		DbMigrate:              true,
		DbLogLevel:             "warning",
	}
}
