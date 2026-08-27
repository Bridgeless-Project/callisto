-- +migrate Up

ALTER TABLE bridge_token_metadata ADD COLUMN dex_name TEXT NOT NULL DEFAULT '';

-- +migrate Down

ALTER TABLE bridge_token_metadata DROP COLUMN dex_name;
