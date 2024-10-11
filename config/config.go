package config

import "github.com/spf13/viper"

type Config struct {
	Server   ServerConfig
	App      AppConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Port int32
}

type AppConfig struct {
	Environment string
	Secret      string
}

type DatabaseConfig struct {
	ConnectionString string
}

func LoadConfig() (*Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	viper.SetDefault("SERVER_PORT", 50051)
	viper.SetDefault("ENVIRONMENT", "Development")

	var cfg Config

	cfg.Server.Port = viper.GetInt32("SERVER_PORT")
	cfg.App.Environment = viper.GetString("DEVELOPMENT")
	cfg.App.Secret = viper.GetString("SECRET_KEY")
	cfg.Database.ConnectionString = viper.GetString("DATABASE_URL")

	return &cfg, nil
}
