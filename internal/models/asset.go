package models

import "time"

type IntervalTime int8

const (
	IntervalTime24h IntervalTime = iota
	IntervalTime7d
	IntervalTime30d
	IntervalTime90d
	IntervalTime1y
	IntervalTimeMax
)

type Asset struct {
	ID         int    `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	Symbol     string `json:"symbol" db:"symbol"`
	SearchName string `json:"-" db:"search_name"`
	ImageURL   string `json:"image_url" db:"image_url"`
}

type Ticker struct {
	PriceUSD     float64   `json:"price_usd" db:"price_usd"`
	MarketCapUSD float64   `json:"marketcap_usd" db:"marketcap_usd"`
	VolumeUSD    float64   `json:"volume_usd" db:"volume_usd"`
	Date         time.Time `json:"date" db:"ticker_date"`
}

type AssetTicker struct {
	Asset
	Ticker
}

type AssetHistory struct {
	Name     string    `json:"asset"`
	Symbol   string    `json:"symbol"`
	Interval string    `json:"interval"`
	Tickers  []*Ticker `json:"tickers"`
}
