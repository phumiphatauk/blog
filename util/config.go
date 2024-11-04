package util

import (
	"os"
	"time"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	Environment             string        `mapstructure:"ENVIRONMENT"`
	DBDriver                string        `mapstructure:"DB_DRIVER"`
	DBSource                string        `mapstructure:"DB_SOURCE"`
	MigrationURL            string        `mapstructure:"MIGRATION_URL"`
	RedisAddress            string        `mapstructure:"REDIS_ADDRESS"`
	HTTPServerAddress       string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	GRPCServerAddress       string        `mapstructure:"GRPC_SERVER_ADDRESS"`
	TokenSymmetricKey       string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration     time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration    time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	URL_FRONTEND            string        `mapstructure:"URL_FRONTEND"`
	GIN_MODE                string        `mapstructure:"GIN_MODE"`
	MINIO_ENDPOINT          string        `mapstructure:"MINIO_ENDPOINT"`
	MINIO_ACCESS_KEY_ID     string        `mapstructure:"MINIO_ACCESS_KEY_ID"`
	MINIO_SECRET_ACCESS_KEY string        `mapstructure:"MINIO_SECRET_ACCESS_KEY"`
	MINIO_USE_SSL           bool          `mapstructure:"MINIO_USE_SSL"`
	MINIO_BUCKET_NAME       string        `mapstructure:"MINIO_BUCKET_NAME"`
	MINIO_URL_RESULT        string        `mapstructure:"MINIO_URL_RESULT"`
	EmailSenderName         string        `mapstructure:"EMAIL_SENDER_NAME"`
	EmailSenderAddress      string        `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EmailSenderPassword     string        `mapstructure:"EMAIL_SENDER_PASSWORD"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	config.Environment = getEnv("ENVIRONMENT", config.Environment)
	config.DBDriver = getEnv("DB_DRIVER", config.DBDriver)
	config.DBSource = getEnv("DB_SOURCE", config.DBSource)
	config.MigrationURL = getEnv("MIGRATION_URL", config.MigrationURL)
	config.RedisAddress = getEnv("REDIS_ADDRESS", config.RedisAddress)
	config.HTTPServerAddress = getEnv("HTTP_SERVER_ADDRESS", config.HTTPServerAddress)
	config.GRPCServerAddress = getEnv("GRPC_SERVER_ADDRESS", config.GRPCServerAddress)
	config.GIN_MODE = getEnv("GIN_MODE", config.GIN_MODE)
	config.URL_FRONTEND = getEnv("URL_FRONTEND", config.URL_FRONTEND)
	config.TokenSymmetricKey = getEnv("TOKEN_SYMMETRIC_KEY", config.TokenSymmetricKey)
	config.MINIO_ENDPOINT = getEnv("MINIO_ENDPOINT", config.MINIO_ENDPOINT)
	config.MINIO_ACCESS_KEY_ID = getEnv("MINIO_ACCESS_KEY_ID", config.MINIO_ACCESS_KEY_ID)
	config.MINIO_SECRET_ACCESS_KEY = getEnv("MINIO_SECRET_ACCESS_KEY", config.MINIO_SECRET_ACCESS_KEY)
	config.MINIO_USE_SSL = getEnvBool("MINIO_USE_SSL", config.MINIO_USE_SSL)
	config.MINIO_BUCKET_NAME = getEnv("MINIO_BUCKET_NAME", config.MINIO_BUCKET_NAME)
	config.MINIO_URL_RESULT = getEnv("MINIO_URL_RESULT", config.MINIO_URL_RESULT)
	config.EmailSenderName = getEnv("EMAIL_SENDER_NAME", config.EmailSenderName)
	config.EmailSenderAddress = getEnv("EMAIL_SENDER_ADDRESS", config.EmailSenderAddress)
	config.EmailSenderPassword = getEnv("EMAIL_SENDER_PASSWORD", config.EmailSenderPassword)

	// accessTokenDuration, _ := strconv.Atoi(os.Getenv("ACCESS_TOKEN_DURATION"))
	// refreshTokenDuration, _ := strconv.Atoi(os.Getenv("REFRESH_TOKEN_DURATION"))

	// config.AccessTokenDuration = time.Duration(accessTokenDuration)
	// config.RefreshTokenDuration = time.Duration(refreshTokenDuration)
	return
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		return value == "true"
	}
	return defaultValue
}
