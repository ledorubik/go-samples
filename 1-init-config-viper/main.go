package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"reflect"
)

type config struct {
	// HTTP server config
	HttpServerRestHost string `mapstructure:"http_server_rest_host"`
	HttpServerRestPort int    `mapstructure:"http_server_rest_port"`
}

func main() {
	// Load config
	cfg, err := initConfig()
	if err != nil {
		log.Fatalf("config init error: %s", err)
	}

	// Print config
	printConfig(cfg)

}

func initConfig() (config *config, err error) {
	viper.AddConfigPath("./1-init-config-viper")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return
}

func printConfig(cfg *config) {
	e := reflect.ValueOf(cfg).Elem()

	for i := 0; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Name
		varType := e.Type().Field(i).Type
		varValue := e.Field(i).Interface()
		fmt.Printf("%v (%v) = %v\n", varName, varType, varValue)
	}
}
