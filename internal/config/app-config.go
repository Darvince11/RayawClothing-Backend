package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	DbUrl      string
	AuthConfig *AuthConfig
}

func Init() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading environmental variables:", err)
	}

	authConfig := AuthConfig{JWTSecretKey: getEnv("JWT_SECRETKEY", "")}

	return &Config{
		Port:       getEnv("PORT", "8080"),
		DbUrl:      getEnv("DATABASE_URL", ""),
		AuthConfig: &authConfig,
	}

}

func getEnv(key, fallback string) string {
	env := os.Getenv(key)
	if env == "" {
		return fallback
	}
	return env
}
