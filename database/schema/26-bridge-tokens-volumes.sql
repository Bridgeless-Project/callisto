-- +migrate Up

CREATE TABLE bridge_tokens_volumes (
    id SERIAL PRIMARY KEY,
    deposit_amount NUMERIC,
    withdrawal_amount NUMERIC,
    commission_amount NUMERIC,
    token_id INTEGER UNIQUE,
    updated_at TIMESTAMP
);


-- +migrate Down

DROP TABLE bridge_tokens_volumes;