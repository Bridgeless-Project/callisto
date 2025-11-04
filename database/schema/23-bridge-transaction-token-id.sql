-- +migrate Up

ALTER TABLE bridge_transactions ADD COLUMN token_id INTEGER NOT NULL DEFAULT 0;

UPDATE bridge_transactions AS bt
SET token_id = bti.token_id
    FROM bridge_tokens_info AS bti
WHERE
    (bt.deposit_chain_id = bti.chain_id
  AND lower(bt.deposit_token) = lower(bti.address)) OR (bt.withdrawal_chain_id = bti.chain_id
  AND lower(bt.withdrawal_token) = lower(bti.address));

-- +migrate Down

ALTER TABLE bridge_transactions DROP COLUMN token_id;