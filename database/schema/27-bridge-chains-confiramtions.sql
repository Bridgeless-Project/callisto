-- +migrate Up

ALTER TABLE bridge_chains ADD COLUMN confirmations INTEGER NOT NULL DEFAULT 0;

-- +migrate Down

ALTER TABLE bridge_chains DROP COLUMN confirmations;
