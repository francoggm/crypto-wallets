package config

import (
	"os"
	"strconv"
)

type Server struct {
	Port string
}

type DB struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

type Tickers struct {
	Interval int
	URL      string
}

type Token struct {
	Expiration int
	SecretKey  string
}

type Config struct {
	Server
	DB
	Tickers
	Token
}

func NewConfig() *Config {
	tiEnv := os.Getenv("TICKERS_INTERVAL_SECONDS")
	ti, err := strconv.Atoi(tiEnv)
	if err != nil {
		// default value = 5 minutes
		ti = 300
	}

	teEnv := os.Getenv("TOKEN_EXPIRATION")
	te, err := strconv.Atoi(teEnv)
	if err != nil {
		// default value = 24 hours
		te = 24
	}

	return &Config{
		Server: Server{
			Port: os.Getenv("SERVER_PORT"),
		},
		DB: DB{
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Name:     os.Getenv("DB_NAME"),
		},
		Tickers: Tickers{
			Interval: ti,
			URL:      os.Getenv("TICKERS_URL"),
		},
		Token: Token{
			Expiration: te,
			SecretKey:  os.Getenv("TOKEN_SECRET_KEY"),
		},
	}
}
