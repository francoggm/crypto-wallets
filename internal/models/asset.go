package models

import "time"

type Asset struct {
	Id         int    `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	Symbol     string `json:"symbol" db:"symbol"`
	SearchName string `json:"-" db:"search_name"`
	ImageURL   string `json:"image_url" db:"image_url"`
}

type AssetTicker struct {
	Asset
	PriceUSD     float64   `json:"price_usd" db:"price_usd"`
	MarketCapUSD float64   `json:"marketcap_usd" db:"marketcap_usd"`
	VolumeUSD    float64   `json:"volume_usd" db:"volume_usd"`
	TickerDate   time.Time `json:"date" db:"ticker_date"`
}
