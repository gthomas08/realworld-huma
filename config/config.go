package config

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
)

type (
	// Config holds the configuration for the application.
	Config struct {
		App      AppConfig      // App holds the configuration related to the application.
		Server   ServerConfig   // Server holds the configuration related to the server.
		Database DatabaseConfig // Database holds the configuration related to the database.
		JWT      JWTConfig      // JWT holds the configuration related to the JWT token.
	}

	// AppConfig holds the configuration related to the application settings.
	AppConfig struct {
		Name    string // Name is the name of the application.
		Version string // Version is the version of the application.
		Env     string // Env is the environment of the application.
	}

	// ServerConfig holds the configuration for the server settings.
	ServerConfig struct {
		Port int // Port is the port to listen on.
	}

	// DatabaseConfig holds the configuration for the database connection.
	DatabaseConfig struct {
		Host     string // Host is the database host.
		Port     int    // Port is the database port.
		Name     string // Name is the name of the database.
		User     string // User is the database user.
		Password string // Password is the database password.
	}

	// JWTConfig holds the configuration for the JWT token.
	JWTConfig struct {
		Key     string // Key is the key used to sign the JWT token.
		Expired int    // Expired is the expiration time of the JWT token.
		Label   string // Label is the label of the JWT token.
	}
)

// LoadConfig loads the configuration from the specified filename.
func LoadConfig(filename string) (*Config, error) {
	// Create a new Viper instance.
	v := viper.New()

	// Set the configuration file name, path, and environment variable settings.
	v.SetConfigName(fmt.Sprintf("config/%s", filename))
	v.AddConfigPath(".")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Read the configuration file.
	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
		return &Config{}, err
	}

	// Unmarshal the configuration into the Config struct.
	var config Config
	if err := v.Unmarshal(&config); err != nil {
		fmt.Printf("Error unmarshaling config: %v\n", err)
		return &Config{}, err
	}

	return &config, nil
}

// LoadConfigPath loads the configuration from the specified path.
func LoadConfigPath(path string) (*Config, error) {
	// Create a new Viper instance.
	v := viper.New()

	// Set the configuration file name, path, and environment variable settings.
	v.SetConfigName(path)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Read the configuration file.
	if err := v.ReadInConfig(); err != nil {
		// Handle the case where the configuration file is not found.
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return &Config{}, errors.New("config file not found")
		}
		return &Config{}, err
	}

	// Parse the configuration into the Config struct.
	var c Config
	if err := v.Unmarshal(&c); err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return &Config{}, err
	}

	return &c, nil
}
