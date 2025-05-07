-- +migrate Up
CREATE TABLE bridge_tokens_info
(
    id SERIAL PRIMARY KEY,
    address VARCHAR(255) NOT NULL UNIQUE ,
    decimals INTEGER NOT NULL,
    chain_id VARCHAR(255) NOT NULL,
    token_id INTEGER NOT NULL,
    is_wrapped BOOLEAN NOT NULL
);

CREATE INDEX idx_bridge_tokens_info_address ON bridge_tokens_info(address);
CREATE INDEX idx_bridge_tokens_info_chain_id ON bridge_tokens_info(chain_id);

CREATE TABLE bridge_token_metadata (
                                token_id VARCHAR(255) UNIQUE PRIMARY KEY,
                                name VARCHAR(255) NOT NULL,
                                symbol VARCHAR(50) NOT NULL,
                                uri TEXT
);

CREATE INDEX idx_bridge_token_metadata_name ON bridge_token_metadata(name);
CREATE INDEX idx_bridge_token_metadata_symbol ON bridge_token_metadata(symbol);

CREATE TABLE bridge_tokens (
                        metadata_id VARCHAR(255) NOT NULL,
                        tokens_info_id BIGINT NOT NULL,
                        commission_rate VARCHAR(255) NOT NULL,
                        FOREIGN KEY (metadata_id) REFERENCES bridge_token_metadata(token_id) ON DELETE CASCADE,
                        FOREIGN KEY (tokens_info_id) REFERENCES bridge_tokens_info(id) ON DELETE CASCADE,
                        PRIMARY KEY (metadata_id, tokens_info_id)
);

CREATE INDEX idx_bridge_tokens_metadata_id ON bridge_tokens(metadata_id);
CREATE INDEX idx_bridge_tokens_tokens_info_id ON bridge_tokens(tokens_info_id);

CREATE TABLE bridge_transactions
(
    id SERIAL PRIMARY KEY,
    deposit_chain_id TEXT NOT NULL,
    deposit_tx_hash TEXT NOT NULL,
    deposit_tx_index BIGINT NOT NULL,
    deposit_block BIGINT NOT NULL,
    deposit_token TEXT NOT NULL,
    deposit_amount BIGINT NOT NULL,
    depositor TEXT NOT NULL,
    receiver TEXT NOT NULL,
    withdrawal_chain_id TEXT,
    withdrawal_tx_hash TEXT,
    withdrawal_token TEXT,
    signature TEXT,
    is_wrapped BOOLEAN NOT NULL,
    withdrawal_amount BIGINT NOT NULL,
    commission_amount BIGINT NOT NULL,
    tx_data TEXT NOT NULL
);

CREATE TABLE bridge_chains
(
    id  TEXT UNIQUE PRIMARY KEY ,
    chain_type SMALLINT,
    bridge_address TEXT,
    operator  TEXT
);

CREATE TABLE transaction_submissions
(
    tx_hash VARCHAR(255) PRIMARY KEY,
    submitters TEXT[]
);
CREATE INDEX idx_tx_hash on transaction_submissions(tx_hash);

CREATE TABLE bridge_params(
    id INT PRIMARY KEY,
    module_admin VARCHAR(255) NOT NULL ,
    parties VARCHAR(255),
    tss_threshold INT
)

-- +migrate Down
DROP TABLE bridge_tokens;
DROP TABLE bridge_tokens_info;
DROP TABLE bridge_chains;
DROP TABLE bridge_transactions;
DROP TABLE bridge_token_metadata;
DROP TABLE transaction_submissions;
DROP TABLE bridge_params;

