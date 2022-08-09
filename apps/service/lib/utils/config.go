package utils

import (
	"fmt"
	"os"
	"strconv"
)

// DefaultConfig defines configuration data that is often used by all services.
type DefaultConfig struct {
	// Environment the service runs on currently.
	Environment Environment

	// IsDebug indicates whether the service is running in debug mode.
	IsDebug bool

	// DatabaseURL is the complete connection URL, using it will overwrite all other
	DatabaseURL      string
	DatabaseUsername string
	DatabasePassword string
	DatabaseHost     string
	DatabasePort     int
	DatabaseName     string

	// HttpAddress is the address where the "normal" HTTP server should listen on.
	HTTPAddress string

	// Version is an arbitrary version string for this instance.
	Version string

	// ServiceName defines how the service is called.
	ServiceName string
	// ServicePrefix defines which prefix is used for getting environment variables.
	// For example "SERVICE" would look for "SERVICE_PORT".
	ServicePrefix string
}

type Environment string

const (
	Development Environment = "development"
	Production  Environment = "production"
)

func GetConfig(servicePrefix string) *DefaultConfig {
	if servicePrefix != "" {
		servicePrefix += "_"
	}

	httpPort := GetEnvOr("PORT", "8001")
	databasePort, err := strconv.Atoi(GetEnvOr(fmt.Sprintf("%sDB_PORT", servicePrefix), "5432"))
	if err != nil {
		panic(err)
	}

	return &DefaultConfig{
		Environment:      GetEnvironment(),
		IsDebug:          IsDebug(),
		DatabaseURL:      GetEnvOr(fmt.Sprintf("%sDB_URL", servicePrefix), ""),
		DatabaseUsername: GetEnvOr(fmt.Sprintf("%sDB_USERNAME", servicePrefix), ""),
		DatabasePassword: GetEnvOr(fmt.Sprintf("%sDB_PASSWORD", servicePrefix), ""),
		DatabaseHost:     GetEnvOr(fmt.Sprintf("%sDB_HOST", servicePrefix), "localhost"),
		DatabasePort:     databasePort,
		DatabaseName:     GetEnvOr(fmt.Sprintf("%sDB_NAME", servicePrefix), "postgres"),
		HTTPAddress:      fmt.Sprintf("0.0.0.0:%s", httpPort),
		Version:          GetEnvOr("VERSION", "1.0"),
		ServicePrefix:    servicePrefix,
	}
}

// IsDebug returns true if DEBUG is set to true.
func IsDebug() bool {
	return GetEnvOr("DEBUG", "true") == "true"
}

// GetEnvironment returns the current environment.
func GetEnvironment() Environment {
	env := GetEnvOr("ENVIRONMENT", "development")

	switch env {
	case "development":
		return Development
	case "production":
		return Production
	default:
		panic("unknown environment")
	}
}

// GetEnvOr returns the given env variable or a default.
func GetEnvOr(key string, otherwise string) string {
	env := os.Getenv(key)

	if env == "" {
		return otherwise
	}
	return env
}

// MustGetEnv returns the given env variable or panics.
func MustGetEnv(key string) string {
	env := os.Getenv(key)

	if env == "" {
		panic(fmt.Sprintf("%s is not set", key))
	}
	return env
}
