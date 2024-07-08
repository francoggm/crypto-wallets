CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  username VARCHAR(100) NOT NULL,
  email VARCHAR(200) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL,
  role VARCHAR(20) NOT NULL DEFAULT 'user' CHECK (role in ('admin', 'user', 'corporate')),
  wallet_usd DOUBLE PRECISION DEFAULT 5000,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT (NOW() AT TIME ZONE 'utc'),
  last_login TIMESTAMP(0) NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc')
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
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT (NOW() AT TIME ZONE 'utc'),

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
  asset_price DECIMAL(40, 25) NOT NULL,
  transaction_date TIMESTAMP NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),

  CONSTRAINT fk_user_wallet FOREIGN KEY (user_id) REFERENCES users(id),
  CONSTRAINT fk_asset_wallet FOREIGN KEY (asset_id) REFERENCES assets(id)
);

CREATE TABLE IF NOT EXISTS tickers (
  id SERIAL PRIMARY KEY,
  asset_id INT NOT NULL,
  price_usd DECIMAL(40, 25) NOT NULL,
  marketcap_usd DECIMAL(40, 25) NOT NULL,
  volume_usd DECIMAL(40, 25) NOT NULL,
  ticker_date TIMESTAMP NOT NULL DEFAULT (NOW() AT TIME ZONE 'utc'),

  CONSTRAINT fk_asset_ticker FOREIGN KEY (asset_id) REFERENCES assets(id)
);

CREATE OR REPLACE FUNCTION get_asset_historical_data(value_id INTEGER, user_interval TEXT)
RETURNS TABLE (
  avg_price_usd DECIMAL(40, 25),
  avg_marketcap_usd DECIMAL(40, 25),
  avg_volume_usd DECIMAL(40, 25),
  trunc_ticker_date TIMESTAMP
) AS $$
DECLARE 
  time_interval INTERVAL;
  trunc_type TEXT;
BEGIN
  trunc_type := 'hour';
  IF user_interval = '24h' THEN
    trunc_type := 'minute';
    time_interval := INTERVAL '1 day';
  ELSIF user_interval = '7d' THEN
    time_interval := INTERVAL '1 week';
  ELSIF user_interval = '30d' THEN
    time_interval := INTERVAL '1 month';
  ELSIF user_interval = '90d' THEN
    time_interval := INTERVAL '3 months';
  ELSIF user_interval = '1y' THEN
    time_interval := INTERVAL '1 year';
  ELSE
    -- Maximum time
    time_interval := INTERVAL '100 years';
  END IF;

  RETURN QUERY 
  SELECT 
    AVG(price_usd) AS avg_price_usd,
    AVG(marketcap_usd) AS avg_marketcap_usd,
    AVG(volume_usd) AS avg_volume_usd,
    date_trunc(trunc_type, ticker_date) AS trunc_ticker_date
  FROM 
    tickers
  WHERE
    asset_id = value_id AND
    ticker_date >= NOW() - time_interval
  GROUP BY 
    date_trunc(trunc_type, ticker_date)
  ORDER BY
    trunc_ticker_date DESC;
END;
$$ LANGUAGE plpgsql;

INSERT INTO assets (name, symbol, search_name, image_url)
VALUES 
('Bitcoin', 'BTC', 'bitcoin', 'https://s2.coinmarketcap.com/static/img/coins/64x64/1.png'),
('Ethereum', 'ETH', 'ethereum', 'https://s2.coinmarketcap.com/static/img/coins/64x64/1027.png'),
('Tether', 'USDT', 'tether', 'https://s2.coinmarketcap.com/static/img/coins/64x64/825.png'),
('Solana', 'SOL', 'solana', 'https://s2.coinmarketcap.com/static/img/coins/64x64/5426.png'),
('XRP', 'XRP', 'xrp', 'https://s2.coinmarketcap.com/static/img/coins/64x64/52.png'),
('Dogecoin', 'DOGE', 'dogecoin', 'https://s2.coinmarketcap.com/static/img/coins/64x64/74.png'),
('Cardano', 'ADA', 'cardano', 'https://s2.coinmarketcap.com/static/img/coins/64x64/2010.png'),
('Shiba Inu', 'SHIB', 'shiba-inu', 'https://s2.coinmarketcap.com/static/img/coins/64x64/5994.png'),
('Avalanche', 'AVAX', 'avalanche', 'https://s2.coinmarketcap.com/static/img/coins/64x64/5994.png'),
('TRON', 'TRX', 'tron', 'https://s2.coinmarketcap.com/static/img/coins/64x64/1958.png');