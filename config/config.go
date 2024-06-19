package config

import (
	"os"
	"strconv"
)

type DB struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

type Config struct {
	DB
	Port            string
	TickersInterval int
	TickersURL      string
}

func NewConfig() *Config {
	tsEnv := os.Getenv("TICKERS_INTERVAL_SECONDS")
	ts, err := strconv.Atoi(tsEnv)
	if err != nil {
		// default value = 5 minutes
		ts = 300
	}

	return &Config{
		Port:            os.Getenv("PORT"),
		TickersInterval: ts,
		TickersURL:      os.Getenv("TICKERS_URL"),
		DB: DB{
			DBUser:     os.Getenv("DBUSER"),
			DBPassword: os.Getenv("DBPASSWORD"),
			DBHost:     os.Getenv("DBHOST"),
			DBPort:     os.Getenv("DBPORT"),
			DBName:     os.Getenv("DBNAME"),
		},
	}
}
