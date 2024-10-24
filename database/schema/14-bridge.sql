-- +migrate Up

CREATE TABLE tokens_info
(
    id SERIAL PRIMARY KEY,
    address  TEXT,
    decimals BIGINT,
    chain_id TEXT,
    token_id INT,
    is_wrapped BOOLEAN
);

CREATE TABLE token_metadata (
    id SERIAL PRIMARY KEY,
    name TEXT,
    symbol TEXT,
    uri TEXT
);

CREATE TABLE tokens (
    id SERIAL PRIMARY KEY,
    metadata_id BIGINT NOT NULL,
    FOREIGN KEY (metadata_id) REFERENCES token_metadata(id) ON DELETE CASCADE
);

CREATE TABLE token_store (
     token_id BIGINT NOT NULL,
     tokens_info_id BIGINT NOT NULL,
     FOREIGN KEY (token_id) REFERENCES tokens(id) ON DELETE CASCADE,
     FOREIGN KEY (tokens_info_id) REFERENCES tokens_info(id) ON DELETE CASCADE,
     PRIMARY KEY (token_id, tokens_info_id)
);

CREATE TABLE transactions
(
    id SERIAL PRIMARY KEY,
    deposit_chain_id TEXT NOT NULL,
    deposit_tx_hash TEXT NOT NULL,
    deposit_tx_index BIGINT NOT NULL,
    deposit_block BIGINT NOT NULL,
    deposit_token TEXT NOT NULL,
    amount BIGINT NOT NULL,
    depositor TEXT NOT NULL,
    receiver TEXT NOT NULL,
    withdrawal_chain_id TEXT,
    withdrawal_tx_hash TEXT,
    withdrawal_token TEXT,
    signature TEXT,
    is_wrapped BOOLEAN NOT NULL
);

CREATE TABLE chains
(
    id  TEXT,
    chain_type SMALLINT,
    bridge_address TEXT,
    operator  TEXT
);

-- +migrate Down
DROP TABLE tokens;
DROP TABLE tokens_info;
DROP TABLE chains;
DROP TABLE transactions;
DROP TABLE token_store;
DROP TABLE token_metadata;

