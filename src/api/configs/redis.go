package configs

import (
	"os"
)

// RedisConfig object
type RedisConfig struct {
	HOST string `env:"REDIS_HOST"`
	PORT string `env:"REDIS_PORT"`
}

// GetRedisConfig returns RedisConfig object
func GetRedisConfig() RedisConfig {
	return RedisConfig{
		HOST: os.Getenv("REDIS_HOST"),
		PORT: os.Getenv("REDIS_PORT"),
	}
}
