CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  username VARCHAR(100) NOT NULL,
  email VARCHAR(200) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS assets (
  id SERIAL PRIMARY KEY,
  name VARCHAR(150) NOT NULL,
  symbol VARCHAR(20) NOT NULL UNIQUE,
  search_name VARCHAR(150) NOT NULL,
  image_url VARCHAR(255)
);

CREATE TABLE IF NOT EXISTS wallets (
  id SERIAL PRIMARY KEY,
  user_id INT NOT NULL,
  asset_id INT NOT NULL,
  quantity DECIMAL(40, 25) NOT NULL DEFAULT 0,

  CONSTRAINT fk_user_wallet FOREIGN KEY (user_id) REFERENCES users(id),
  CONSTRAINT fk_asset_wallet FOREIGN KEY (asset_id) REFERENCES assets(id),
  UNIQUE(user_id, asset_id)
);

CREATE TABLE IF NOT EXISTS transactions (
  id SERIAL PRIMARY KEY,
  user_id INT NOT NULL,
  asset_id INT NOT NULL,
  transaction_type VARCHAR(10) NOT NULL CHECK (transaction_type in ('buy', 'sell')),
  amount DECIMAL(40, 25) NOT NULL,
  transaction_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

  CONSTRAINT fk_user_wallet FOREIGN KEY (user_id) REFERENCES users(id),
  CONSTRAINT fk_asset_wallet FOREIGN KEY (asset_id) REFERENCES assets(id)
);

CREATE TABLE IF NOT EXISTS tickers (
  id SERIAL PRIMARY KEY,
  asset_id INT NOT NULL,
  price_usd DECIMAL(40, 25) NOT NULL,
  marketcap_usd DECIMAL(40, 25) NOT NULL,
  volume_usd DECIMAL(40, 25) NOT NULL,
  ticker_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

  CONSTRAINT fk_asset_ticker FOREIGN KEY (asset_id) REFERENCES assets(id)
);

INSERT INTO assets (name, symbol, search_name, image_url)
VALUES 
('Bitcoin', 'BTC', 'bitcoin', 'https://s2.coinmarketcap.com/static/img/coins/64x64/1.png'),
('Ethereum', 'ETH', 'ethereum', 'https://s2.coinmarketcap.com/static/img/coins/64x64/1027.png'),
('Tether', 'USDT', 'tether', 'https://s2.coinmarketcap.com/static/img/coins/64x64/825.png'),
('Solana', 'SOL', 'solana', 'https://s2.coinmarketcap.com/static/img/coins/64x64/5426.png'),
('XRP', 'XRP', 'ripple', 'https://s2.coinmarketcap.com/static/img/coins/64x64/52.png'),
('Dogecoin', 'DOGE', 'dogecoin', 'https://s2.coinmarketcap.com/static/img/coins/64x64/74.png'),
('Cardano', 'ADA', 'cardano', 'https://s2.coinmarketcap.com/static/img/coins/64x64/2010.png'),
('Shiba Inu', 'SHIB', 'shiba-inu', 'https://s2.coinmarketcap.com/static/img/coins/64x64/5994.png'),
('Avalanche', 'AVAX', 'avalanche', 'https://s2.coinmarketcap.com/static/img/coins/64x64/5994.png'),
('TRON', 'TRX', 'tron', 'https://s2.coinmarketcap.com/static/img/coins/64x64/1958.png');