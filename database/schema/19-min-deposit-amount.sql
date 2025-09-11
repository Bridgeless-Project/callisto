-- +migrate Up
ALTER TABLE bridge_tokens_info ADD COLUMN min_deposit TEXT DEFAULT "";

-- +migrate Down
ALTER TABLE bridge_tokens_info DROP COLUMN min_deposit;