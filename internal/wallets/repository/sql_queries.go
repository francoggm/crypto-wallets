package repository

const (
	getUserWallets = "SELECT * FROM wallets WHERE user_id = $1"

	getUserWallet = "SELECT * FROM wallets WHERE user_id = $1 AND asset_id = $2"

	insertWallet = "INSERT INTO wallets (user_id, asset_id) VALUES (:user_id, :asset_id) RETURNING *"

	updateWallet = "UPDATE wallets SET quantity = :quantity, updated_at = :updated_at WHERE user_id = :user_id AND asset_id = :asset_id RETURNING *"

	updateWalletUSD = "UPDATE users SET wallet_usd = $2, updated_at = $3 WHERE id = $1"

	insertTransaction = "INSERT INTO transactions (user_id, asset_id, transaction_type, amount, asset_price) VALUES ($1, $2, $3, $4, $5)"
)
