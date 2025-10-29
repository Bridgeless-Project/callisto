-- +migrate Up
CREATE DOMAIN uint16 AS integer
    CHECK (VALUE BETWEEN 0 AND 65535);


CREATE TABLE referral
(
    id uint16 PRIMARY KEY,
    withdrawal_address TEXT,
    commission_rate INT
);

CREATE TABLE referral_rewards
(
    referral_id          uint16,
    token_id             BIGINT,
    to_claim             string,
    total_claimed_amount string
);

ALTER TABLE bridge_transactions ADD COLUMN uint16 INT DEFAULT 0;

-- +migrate Down
DROP TABLE referral;
DROP TABLE referral_rewards;

ALTER TABLE bridge_transactions DROP COLUMN referral_id;

DROP DOMAIN uint16;