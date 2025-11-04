-- +migrate Up

ALTER TABLE bridge_transactions ADD COLUMN withdrawal_decimals INTEGER NOT NULL DEFAULT 0;

UPDATE bridge_transactions AS bt
SET withdrawal_decimals = bti.decimals
    FROM bridge_tokens_info AS bti
WHERE
    bt.withdrawal_chain_id = bti.chain_id
  AND lower(bt.withdrawal_token) = lower(bti.address);

-- +migrate Down

ALTER TABLE bridge_transactions DROP COLUMN withdrawal_decimals;