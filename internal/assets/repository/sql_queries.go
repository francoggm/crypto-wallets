package repository

const (
	getAllAssets = "SELECT * FROM assets"

	getAsset = "SELECT * FROM assets WHERE name = $1 OR search_name = $2 OR symbol = $3"

	getAssetTicker = `
		SELECT price_usd, marketcap_usd, volume_usd, ticker_date 
		FROM tickers WHERE asset_id = $1 
		ORDER BY ticker_date DESC 
		LIMIT 1`
	
	saveAssetTicker = "INSERT INTO tickers (asset_id, marketcap_usd, volume_usd, price_usd) VALUES (:id, :marketcap_usd, :volume_usd, :price_usd)"
)
