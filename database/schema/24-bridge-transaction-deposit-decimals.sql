-- +migrate Up

ALTER TABLE bridge_transactions ADD COLUMN deposit_decimals INTEGER NOT NULL DEFAULT 0;

UPDATE bridge_transactions AS bt
SET deposit_decimals = bti.decimals
    FROM bridge_tokens_info AS bti
WHERE
    bt.deposit_chain_id = bti.chain_id
  AND lower(bt.deposit_token) = lower(bti.address);

-- +migrate Down

ALTER TABLE bridge_transactions DROP COLUMN deposit_decimals;