package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

type Config struct {
	Database DatabaseConfig `mapstructure:"database"`
}

type DatabaseConfig struct {
	BaseUrl   string `mapstructure:"baseurl"`
	URL       string `mapstructure:"-"`
	MaxConns  int    `mapstructure:"max_conns"`
	MinConns  int    `mapstructure:"min_conns"`
	User      string `mapstructure:"-"`
	Password  string `mapstructure:"-"`
	Name      string `mapstructure:"-"`
	Container string `mapstructure:"-"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("finance-config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	config.Database.User = os.Getenv("POSTGRES_USER")
	config.Database.Password = os.Getenv("POSTGRES_PASSWORD")
	config.Database.Name = os.Getenv("POSTGRES_DB")
	config.Database.Container = os.Getenv("POSTGRES_CONTAINER")

	if config.Database.User == "" || config.Database.Password == "" || config.Database.Name == "" || config.Database.Container == "" {
		log.Fatal("Database credentials and settings are not fully set in environment variables")
	}

	config.Database.URL = fmt.Sprintf("%s%s?sslmode=disable", config.Database.BaseUrl, config.Database.Name)

	return &config, nil
}

func MustLoad() (*Config, error) {
	config, err := LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	return config, nil
}
