-- +migrate Up

ALTER TABLE bridge_params
    RENAME COLUMN relayer_account TO relayer_accounts;

ALTER TABLE bridge_params
    ALTER COLUMN relayer_accounts DROP DEFAULT;

ALTER TABLE bridge_params
ALTER COLUMN relayer_accounts TYPE TEXT[]
    USING ARRAY[relayer_accounts];

ALTER TABLE bridge_params
    ALTER COLUMN relayer_accounts SET DEFAULT '{}';

-- +migrate Down

ALTER TABLE bridge_params
    ALTER COLUMN relayer_accounts DROP DEFAULT;

ALTER TABLE bridge_params
ALTER COLUMN relayer_accounts TYPE VARCHAR(255)
        USING relayer_accounts[1];

ALTER TABLE bridge_params
    RENAME COLUMN relayer_accounts TO relayer_account;

ALTER TABLE bridge_params
    ALTER COLUMN relayer_account SET DEFAULT '';