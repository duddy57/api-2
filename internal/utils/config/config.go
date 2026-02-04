package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Port         string
	Environment  string
	LogLevel     string
	Database     DatabaseConfig
	Server       ServerConfig
	Redis        RedisConfig
	ResendAPIKey string
}

type RedisConfig struct {
	Host         string
	User         string
	Port         int
	Password     string
	DB           int
	PoolSize     int
	MinIdleConns int
	MaxConnAge   time.Duration
}

type DatabaseConfig struct {
	Host           string
	Port           string
	User           string
	Password       string
	Name           string
	SSLMode        string
	ChannelBinding string
	MaxOpenConns   int
	MaxIdleConns   int
	MaxLifetime    time.Duration

	GooseString   string
	MigrationsDir string
}

type ServerConfig struct {
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
	ServerUrl    string
	FrontendUrl  string
}

func Load() *Config {
	return &Config{
		Port:        getEnv("PORT", "8080"),
		Environment: getEnv("ENVIRONMENT", "development"),
		LogLevel:    getEnv("LOG_LEVEL", "info"),
		Database: DatabaseConfig{
			Host:           getEnv("DB_HOST", "localhost"),
			Port:           getEnv("DB_PORT", "5432"),
			User:           getEnv("DB_USER", "postgres"),
			Password:       getEnv("DB_PASSWORD", ""),
			Name:           getEnv("DB_NAME", "api"),
			SSLMode:        getEnv("DB_SSL_MODE", "disable"),
			ChannelBinding: getEnv("DB_CHANNELBINDING", ""),
			MaxOpenConns:   getEnvAsInt("DB_MAX_OPEN_CONNS", 25),
			MaxIdleConns:   getEnvAsInt("DB_MAX_IDLE_CONNS", 25),
			MaxLifetime:    getEnvAsDuration("DB_MAX_LIFETIME", "5m"),
			GooseString:    getEnv("GOOSE_STRING", ""),
			MigrationsDir:  getEnv("GOOSE_MIGRATIONS_DIR", "./internal/store/pgstore/migrations"),
		},
		Server: ServerConfig{
			ReadTimeout:  getEnvAsDuration("SERVER_READ_TIMEOUT", "10s"),
			WriteTimeout: getEnvAsDuration("SERVER_WRITE_TIMEOUT", "10s"),
			IdleTimeout:  getEnvAsDuration("SERVER_IDLE_TIMEOUT", "60s"),
			ServerUrl:    getEnv("SERVER_URL", "http://localhost:8080"),
			FrontendUrl:  getEnv("FRONTEND_URL", "http://localhost:3000"),
		},
		Redis: RedisConfig{
			Host:         getEnv("REDIS_HOST", "localhost"),
			User:         getEnv("REDIS_USER", ""),
			Port:         getEnvAsInt("REDIS_PORT", 6379),
			Password:     getEnv("REDIS_PASSWORD", ""),
			DB:           getEnvAsInt("REDIS_DB", 0),
			PoolSize:     getEnvAsInt("REDIS_POOL_SIZE", 10),
			MinIdleConns: getEnvAsInt("REDIS_MIN_IDLE_CONNS", 2),
			MaxConnAge:   getEnvAsDuration("REDIS_MAX_CONN_AGE", "5m"),
		},
		ResendAPIKey: getEnv("RESEND_API_KEY", ""),
	}
}

func (c *Config) GetDatabaseURL() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Database.Host,
		c.Database.Port,
		c.Database.User,
		c.Database.Password,
		c.Database.Name,
		c.Database.SSLMode,
	)
}

func (c *Config) GetRedisURL() string {
	return fmt.Sprintf(
		"redis://%s:%s@%s:%d/%d",
		c.Redis.User,
		c.Redis.Password,
		c.Redis.Host,
		c.Redis.Port,
		c.Redis.DB,
	)
}

func (c *Config) GooseDSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Database.Host,
		c.Database.Port,
		c.Database.User,
		c.Database.Password,
		c.Database.Name,
		c.Database.SSLMode,
	)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
		log.Printf("Warning: Invalid integer value for %s: %s, using default: %d", key, value, defaultValue)
	}
	return defaultValue
}

func getEnvAsDuration(key string, defaultValue string) time.Duration {
	if value := os.Getenv(key); value != "" {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
		log.Printf("Warning: Invalid duration value for %s: %s, using default: %s", key, value, defaultValue)
	}
	duration, _ := time.ParseDuration(defaultValue)
	return duration
}
