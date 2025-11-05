-- +migrate Up

CREATE TABLE bridge_tokens_volumes (
    id SERIAL PRIMARY KEY,
    deposit_amount NUMERIC,
    withdrawal_amount NUMERIC,
    commission_amount NUMERIC,
    token_id INTEGER,
    updated_at TIMESTAMP
);

CREATE UNIQUE INDEX uniq_tokens_volume ON bridge_tokens_volumes(token_id,updated_at);

-- +migrate Down

DROP TABLE bridge_tokens_volumes;
DROP INDEX uniq_tokens_volume;