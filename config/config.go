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
	BaseUrl   string `mapstructure:"base_url"`
	Driver    string `mapstructure:"driver"`
	MaxConns  int    `mapstructure:"max_conns"`
	MinConns  int    `mapstructure:"min_conns"`
	URL       string `mapstructure:"-"`
	User      string `mapstructure:"-"`
	Password  string `mapstructure:"-"`
	Name      string `mapstructure:"-"`
	Container string `mapstructure:"-"`
}

func LoadConfig() (*Config, error) {
	configPaths := []string{
		".",     // current directory
		"..",    // parent directory
		"../..", // two levels up
		"/home/runner/work/FinanceApp/FinanceApp/config",
		os.Getenv("CONFIG_PATH"), // environment variable to specify path
	}

	viper.SetConfigName("finance-config")
	viper.SetConfigType("yaml")

	for _, path := range configPaths {
		viper.AddConfigPath(path) // adds each path to the search list
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	db := config.Database
	db.User = os.Getenv("POSTGRES_USER")
	db.Password = os.Getenv("POSTGRES_PASSWORD")
	db.Name = os.Getenv("POSTGRES_DB")
	db.Container = os.Getenv("POSTGRES_CONTAINER")

	if db.User == "" || db.Password == "" || db.Name == "" || db.Container == "" {
		log.Fatalf("user: %s  password: %s db.name: %s db.container: %s", db.User, db.Password, db.Name, db.Container)
		log.Fatal("Database credentials and settings are not fully set in environment variables")
	}

	config.Database.URL = fmt.Sprintf("%s://%s:%s%s%s?sslmode=disable", db.Driver, db.User, db.Password, db.BaseUrl, db.Name)

	return &config, nil
}

func MustLoad() (*Config, error) {
	config, err := LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	return config, nil
}
