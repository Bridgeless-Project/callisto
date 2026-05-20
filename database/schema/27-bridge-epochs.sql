-- +migrate Up
ALTER TABLE bridge_params ADD COLUMN IF NOT EXISTS relayer_accounts TEXT[] DEFAULT '{}';
ALTER TABLE bridge_params ADD COLUMN IF NOT EXISTS epoch_id INTEGER DEFAULT 0;
ALTER TABLE bridge_params ADD COLUMN IF NOT EXISTS supporting_time BIGINT DEFAULT 0;

CREATE TABLE bridge_epochs
(
    id INTEGER PRIMARY KEY,
    status SMALLINT NOT NULL,
    start_block BIGINT NOT NULL DEFAULT 0,
    end_block BIGINT NOT NULL DEFAULT 0,
    parties TEXT[],
    tss_threshold INTEGER NOT NULL DEFAULT 0,
    tss_info JSONB NOT NULL DEFAULT '[]'::jsonb,
    finalized_block BIGINT NOT NULL DEFAULT 0
);

CREATE TABLE bridge_epoch_chain_signatures
(
    epoch_id INTEGER NOT NULL,
    chain_type SMALLINT NOT NULL,
    added_mod SMALLINT,
    added_epoch_id INTEGER,
    added_signature TEXT,
    added_new_signer TEXT,
    added_start_time BIGINT,
    added_end_time BIGINT,
    added_nonce TEXT,
    removed_mod SMALLINT,
    removed_epoch_id INTEGER,
    removed_signature TEXT,
    removed_new_signer TEXT,
    removed_start_time BIGINT,
    removed_end_time BIGINT,
    removed_nonce TEXT,
    PRIMARY KEY (epoch_id, chain_type)
);

CREATE TABLE bridge_epoch_pubkeys
(
    epoch_id INTEGER PRIMARY KEY,
    pubkey TEXT NOT NULL
);

CREATE TABLE bridge_epoch_pubkey_submissions
(
    epoch_id INTEGER NOT NULL,
    hash TEXT NOT NULL,
    submitters TEXT[],
    PRIMARY KEY (epoch_id, hash)
);

CREATE TABLE bridge_epoch_signature_submissions
(
    epoch_id INTEGER NOT NULL,
    hash TEXT NOT NULL,
    submitters TEXT[],
    PRIMARY KEY (epoch_id, hash)
);

-- +migrate Down
DROP TABLE bridge_epoch_signature_submissions;
DROP TABLE bridge_epoch_pubkey_submissions;
DROP TABLE bridge_epoch_pubkeys;
DROP TABLE bridge_epoch_chain_signatures;
DROP TABLE bridge_epochs;
ALTER TABLE bridge_params DROP COLUMN IF EXISTS supporting_time;
ALTER TABLE bridge_params DROP COLUMN IF EXISTS epoch_id;
ALTER TABLE bridge_params DROP COLUMN IF EXISTS relayer_accounts;
