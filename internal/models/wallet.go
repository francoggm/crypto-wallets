package models

import "time"

type Wallet struct {
	ID        int       `json:"id" db:"id"`
	UserID    int       `json:"user_id" db:"user_id"`
	AssetID   int       `json:"asset_id" db:"asset_id"`
	Quantity  float64   `json:"quantity" db:"quantity"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type WalletTicker struct {
	Symbol    string    `json:"symbol"`
	Asset     string    `json:"asset"`
	PriceUSD  float64   `json:"price_usd"`
	Quantity  float64   `json:"quantity"`
	TotalUSD  float64   `json:"total_usd"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TransactionAsset struct {
	Asset    string  `json:"asset"`
	Quantity float64 `json:"quantity"`
}
