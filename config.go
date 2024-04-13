package main

import (
	"fmt"
	"os"
)

// env vars

type Config struct {
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
	JWTSecret  string
}

// config for sql
func initConfig() Config {
	return Config{
		Port:       getEnv("PORT", ":8080"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWD", "password"),
		DBAddress:  fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName:     getEnv("DB_NAME", "projectmanager"),
		JWTSecret:  getEnv("JWT_SECRET", "randomjwtsecretkey"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

var Envs = initConfig()
