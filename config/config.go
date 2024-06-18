package config

import "os"

type Config struct {
	Port       string
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

func NewConfig() *Config {
	return &Config{
		Port:       os.Getenv("PORT"),
		DBUser:     os.Getenv("DBUSER"),
		DBPassword: os.Getenv("DBPASSWORD"),
		DBHost:     os.Getenv("DBHOST"),
		DBPort:     os.Getenv("DBPORT"),
		DBName:     os.Getenv("DBNAME"),
	}
}
