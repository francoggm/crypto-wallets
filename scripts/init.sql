CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  username VARCHAR(100) NOT NULL,
  email VARCHAR(200) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS assets (
  id SERIAL PRIMARY KEY,
  name VARCHAR(150) NOT NULL,
  symbol VARCHAR(20) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS wallets (
  id SERIAL PRIMARY KEY,
  user_id INT NOT NULL,
  asset_id INT NOT NULL,
  quantity DECIMAL(18, 8) NOT NULL DEFAULT 0,

  CONSTRAINT fk_user_wallet FOREIGN KEY (user_id) REFERENCES users(id),
  CONSTRAINT fk_asset_wallet FOREIGN KEY (asset_id) REFERENCES assets(id),
  UNIQUE(user_id, asset_id)
);

CREATE TABLE IF NOT EXISTS transactions (
  id SERIAL PRIMARY KEY,
  user_id INT NOT NULL,
  asset_id INT NOT NULL,
  transaction_type VARCHAR(10) NOT NULL CHECK (transaction_type in ('buy', 'sell')),
  amount DECIMAL(18, 8) NOT NULL,
  transaction_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

  CONSTRAINT fk_user_wallet FOREIGN KEY (user_id) REFERENCES users(id),
  CONSTRAINT fk_asset_wallet FOREIGN KEY (asset_id) REFERENCES assets(id)
);

CREATE TABLE IF NOT EXISTS tickers (
  id SERIAL PRIMARY KEY,
  asset_id INT NOT NULL,
  price_usd DECIMAL(18, 8) NOT NULL,
  marketcap_usd DECIMAL(18, 8) NOT NULL,
  volume_usd DECIMAL(18, 8) NOT NULL,
  ticker_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

  CONSTRAINT fk_asset_ticker FOREIGN KEY (asset_id) REFERENCES assets(id)
);