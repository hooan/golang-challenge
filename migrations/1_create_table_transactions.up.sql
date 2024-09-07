CREATE TABLE IF NOT EXISTS transactions(
    id BIGSERIAL NOT NULL PRIMARY KEY,
    tx_id INTEGER NOT NULL,
    tx_date VARCHAR(255),
    amount FLOAT
)