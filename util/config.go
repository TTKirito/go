package util

import (
	"time"

	"github.com/spf13/viper"
)

// Config store all configuration of the application
// The values are read by viper from a config file or environment variables.

type Config struct {
	DBDriver                string        `mapstructure:"DB_DRIVER"`
	DBSOURCE                string        `mapstructure:"DB_SOURCE"`
	MigrationURL            string        `mapstructure:"MIGRATION_URL"`
	HTTPServerAddress       string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	GRPCServerAddress       string        `mapstructure:"GRPC_SERVER_ADDRESS"`
	RedisAddress            string        `mapstructure:"REDIS_ADDRESS"`
	TokenSymmetricKey       string        `mapstructure:"TOKEN_SYMETRIC_KEY"`
	AccessTokenDuration     time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefrefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	EmailSenderName         string        `mapstructure:"EMAIL_SENDER_NAME"`
	EmailSenderAddress      string        `mapstructure:"EMAIL_SENDER_ADDRESS"`
	EmailSenderPassword     string        `mapstructure:"EMAIL_SENDER_PASSWORD"`
}

// LoadConfig reads configuration from file or enviroment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env") // json, xml

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
