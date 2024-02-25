package config

type Config struct {
	// HTTP server config
	HttpServerGinMode                 string `mapstructure:"http_server_gin_mode"`
	HttpServerTlsEnable               bool   `mapstructure:"http_server_tls_enable"`
	HttpServerTlsServerCertPath       string `mapstructure:"http_server_tls_server_cert_path"`
	HttpServerTlsServerKeyPath        string `mapstructure:"http_server_tls_server_key_path"`
	HttpServerTlsSkipServerCertVerify bool   `mapstructure:"http_server_tls_skip_server_cert_verify"`
	HttpServerRestHost                string `mapstructure:"http_server_rest_host"`
	HttpServerRestPort                int    `mapstructure:"http_server_rest_port"`
}

func New() *Config {
	return &Config{
		HttpServerGinMode:                 "release",
		HttpServerTlsEnable:               false,
		HttpServerTlsServerCertPath:       "",
		HttpServerTlsServerKeyPath:        "",
		HttpServerTlsSkipServerCertVerify: false,
		HttpServerRestHost:                "localhost",
		HttpServerRestPort:                8080,
	}
}
