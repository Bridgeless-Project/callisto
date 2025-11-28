-- +migrate Up

ALTER TABLE bridge_transactions ADD COLUMN referral_id INTEGER NOT NULL DEFAULT 0;

-- +migrate Down

ALTER TABLE bridge_transactions DROP COLUMN referral_id;