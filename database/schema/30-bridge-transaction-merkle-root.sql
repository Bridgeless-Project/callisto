-- +migrate Up

ALTER TABLE bridge_transactions ADD COLUMN merkle_root TEXT NOT NULL DEFAULT '';

-- +migrate Down

ALTER TABLE bridge_transactions DROP COLUMN merkle_root;