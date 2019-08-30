package ankrmicro

import (
	"fmt"
	"os"
	"strings"
)

// Config struct save the configuration of the micro-service
// Right now, the service is only using rabbitmq and monge as
// external dependency.
type Config struct {
	Rabbitmq     string
	DatabaseHost string
	DatabaseName string
	Listen       string
	DevEnv       bool
}

var config Config

func init() {
	config.DatabaseHost = "localhost:27017"
	config.Rabbitmq = "amqp://guest:guest@localhost:5672"
	config.Listen = ":50051"
	config.DatabaseName = "dccn"
	config.DevEnv = true
	config = LoadConfigFromEnv()

}

// GetConfig function reads configurations from specified source and
// return as a Config struct. It could be expended to new sources in
// the future.
func GetConfig() Config {
	return LoadConfigFromEnv()
}

// LoadConfigFromEnv function reads configurations from environment
// variables.
func LoadConfigFromEnv() Config {
	value := os.Getenv("MICRO_BROKER_ADDRESS")
	if len(value) > 0 {
		config.Rabbitmq = value
	}

	value = os.Getenv("DB_HOST")
	if len(value) > 0 {
		config.DatabaseHost = value
	}

	value = os.Getenv("DB_NAME")

	if len(value) > 0 {
		config.DatabaseName = value
	}

	value = os.Getenv("MICRO_SERVER_ADDRESS")

	if len(value) > 0 {
		config.Listen = value
	}


	value = os.Getenv("DEV_EVN")

	if len(value) > 0 {
		value = strings.ToLower(value)
		if strings.Compare(value, "false") == 0 {
			config.DevEnv = false
		}

	}

	return config
}

// Show function print the configuration information to
// the standard output.
func (config *Config) Show() {
	fmt.Printf("RabbitMQ : %s \n", config.Rabbitmq)
	fmt.Printf("DB_HOST  : %s  \n", config.DatabaseHost)
	fmt.Printf("DB_Name  : %s  \n", config.DatabaseName)
	fmt.Printf("Listen   : %s \n", config.Listen)
	fmt.Printf("DevEnv   : %t \n", config.DevEnv)
}
