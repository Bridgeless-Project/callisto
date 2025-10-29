-- +migrate Up
CREATE DOMAIN uint16 AS integer
    CHECK (VALUE BETWEEN 0 AND 65535);


CREATE TABLE bridge_referral
(
    id uint16 PRIMARY KEY,
    withdrawal_address TEXT,
    commission_rate INT
);

CREATE TABLE bridge_referral_rewards
(
    referral_id          uint16,
    token_id             BIGINT,
    to_claim             TEXT,
    total_claimed_amount TEXT
);

ALTER TABLE bridge_transactions ADD COLUMN referral_id uint16 DEFAULT 0;

-- +migrate Down
DROP TABLE bridge_referral;
DROP TABLE bridge_referral_rewards;

ALTER TABLE bridge_transactions DROP COLUMN referral_id;

DROP DOMAIN uint16;