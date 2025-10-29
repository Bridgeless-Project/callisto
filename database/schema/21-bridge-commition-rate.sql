-- +migrate Up
ALTER TABLE bridge_tokens_info ADD COLUMN commission_rate VARCHAR(255) NOT NULL DEFAULT '';

UPDATE bridge_tokens_info
SET commission_rate = bridge_tokens.commission_rate
    FROM bridge_tokens
WHERE bridge_tokens_info.id = bridge_tokens.tokens_info_id;

ALTER TABLE bridge_tokens DROP COLUMN commission_rate;

-- +migrate Down
ALTER TABLE bridge_tokens ADD COLUMN commission_rate VARCHAR(255) NOT NULL DEFAULT '';

UPDATE bridge_tokens
SET commission_rate = bridge_tokens_info.commission_rate
    FROM bridge_tokens_info
WHERE bridge_tokens.tokens_info_id = bridge_tokens_info.id;

ALTER TABLE bridge_tokens_info DROP COLUMN commission_rate;


