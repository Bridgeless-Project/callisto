-- +migrate Up

ALTER TABLE bridge_tranasctions ADD COLUMN withdrawal_tx_hash TEXT NOT NULL DEFAULT '';

-- +migrate Down

ALTER TABLE bridge_transactions DROP COLUMN withdrawal_tx_hash;