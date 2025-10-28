-- +migrate Up
ALTER TABLE bridge_transactions ADD COLUMN core_tx_timestamp TIMESTAMP;

-- +migrate Down
ALTER TABLE bridge_transactions DROP COLUMN core_tx_timestamp;