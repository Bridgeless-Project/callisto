-- +migrate Up
ALTER TABLE bridge_tokens_info ADD COLUMN commission_rate VARCHAR(255) NOT NULL;

UPDATE bridge_tokens_info
    JOIN bridge_tokens ON bridge_tokens_info.id = bridge_tokens.tokens_info_id
    SET bridge_tokens_info.commission_rate = bridge_tokens.commission_rate;

ALTER TABLE bridge_tokens DROP COLUMN commission_rate;

-- +migrate Down

ALTER TABLE bridge_tokens ADD COLUMN commission_rate VARCHAR(255) NOT NULL;

UPDATE bridge_tokens
    JOIN bridge_tokens ON bridge_tokens.tokens_info_id = bridge_tokens_info.id
    SET bridge_tokens.commission_rate = bridge_tokens_info.commission_rate;

ALTER TABLE bridge_tokens_info DROP COLUMN commission_rate;


