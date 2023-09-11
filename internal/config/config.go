package config

import (
	"github.com/spf13/viper"
)

// AppConfig contains application configuration settings.
type AppConfig struct {
	Port       int
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	// Add more configuration fields as needed
}

// LoadConfig loads application configuration from environment variables and/or a configuration file.
func LoadConfig() (*AppConfig, error) {
	// Initialize Viper

	viper.AutomaticEnv() // Read environment variables

	// Read configuration from file (optional)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	// Define default configuration values
	viper.SetDefault("Port", 8080)
	viper.SetDefault("DBHost", "localhost")
	viper.SetDefault("DBPort", 5432)
	viper.SetDefault("DBUser", "username")
	viper.SetDefault("DBPassword", "password")
	viper.SetDefault("DBName", "mydb")

	// Unmarshal configuration into the struct
	var cfg AppConfig
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
