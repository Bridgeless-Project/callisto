-- +migrate Up

CREATE TABLE bridge_tokens_volumes (
    id SERIAL PRIMARY KEY,
    deposit_amount BIGINT,
    withdrawal_amount BIGINT,
    commission_amount BIGINT,
    token_id INTEGER UNIQUE,
    update_timestamp TIMESTAMP
);

CREATE UNIQUE INDEX uniq_token_volume ON bridge_tokens_volumes(token_id,update_timestamp);

-- +migrate Down

DROP TABLE bridge_tokens_volumes;
DROP UNIQUE INDEX uniq_token_volume;
