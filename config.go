package main

import "os"

type Config struct {
	Port string

	DBName    string
	JWTSecret string
}

var Envs = initConfig()

func initConfig() Config {
	return Config{
		Port:      getEnv("PORT", "8080"),
		DBName:    getEnv("DB_NAME", "projectmanager"),
		JWTSecret: getEnv("JWT_SECRET", "randomjwtsecretkey"),
	}
}
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value

	}
	return fallback
}
