-- +migrate Up
CREATE TABLE referral {
    id INT PRIMARY KEY,
    withdrawal_address TEXT,
    commission_rate INT
};

CREATE TABLE referral_rewards {
    referral_id UINT,
    token_id BIGINT,
    to_claim COIN,
    total_collected_amount COIN,
};

ALTER TABLE bridge_transactions ADD COLUMN referral_id INT DEFAULT 0;

-- +migrate Down
DROP TABLE referral;
DROP TABLE referral_rewards;

ALTER TABLE bridge_transactions DROP COLUMN referral_id;