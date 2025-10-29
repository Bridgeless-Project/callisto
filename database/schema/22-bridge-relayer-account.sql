-- +migrate Up

ALTER TABLE bridge_params ADD COLUMN relayer_account VARCHAR(255) NOT NULL;

-- +migrate Down

ALTER TABLE bridge_params DROP COLUMN relayer_account;
