package assets

import "errors"

var (
	InvalidAsset                = "Invalid asset!"
	AssetNotFound               = "Asset not found!"
	FailedGettingTicker         = "Failed getting ticker!"
	FailedGettingHistoricalData = "Failed getting historical data!"
)

var (
	ErrInvalidAsset                = errors.New(InvalidAsset)
	ErrAssetNotFound               = errors.New(AssetNotFound)
	ErrFailedGettingTicker         = errors.New(FailedGettingTicker)
	ErrFailedGettingHistoricalData = errors.New(FailedGettingHistoricalData)
)
