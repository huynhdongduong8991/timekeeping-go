// Setup viper configuration

package lib

import (
	"github.com/spf13/viper"
)

// Contract struct
type Contract struct {
	HTTP_RPC_URL     string
	WS_RPC_URL       string
	CONTRACT_ADDRESS string
	FROM_BLOCK       int64
}

// DB struct
type DB struct {
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	SSL_MODE    string
}

// Config struct
type Config struct {
	Contract Contract
	DB       DB
}

// NewConfig function
func NewConfig() (*Config, error) {
	viper.AddConfigPath("../")
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return &Config{
		Contract: Contract{
			HTTP_RPC_URL:     viper.GetString("HTTP_RPC_URL"),
			WS_RPC_URL:       viper.GetString("WS_RPC_URL"),
			CONTRACT_ADDRESS: viper.GetString("CONTRACT_ADDRESS"),
			FROM_BLOCK:       viper.GetInt64("FROM_BLOCK"),
		},
		DB: DB{
			DB_HOST:     viper.GetString("DB_HOST"),
			DB_PORT:     viper.GetString("DB_PORT"),
			DB_USER:     viper.GetString("DB_USER"),
			DB_PASSWORD: viper.GetString("DB_PASSWORD"),
			DB_NAME:     viper.GetString("DB_NAME"),
			SSL_MODE:    viper.GetString("SSL_MODE"),
		},
	}, nil
}
