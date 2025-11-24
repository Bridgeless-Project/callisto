-- +migrate Up

ALTER TABLE bridge_chains ADD COLUMN name TEXT NOT NULL DEFAULT '';

-- +migrate Down

ALTER TABLE bridge_chains DROP COLUMN name;
