-- +migrate Up
CREATE TABLE admins_vesting
(
    id SERIAL,
    address  TEXT,
    vesting_period INTEGER,
    reward_per_period  COIN[] NOT NULL DEFAULT '{}',
    last_vesting_time  INTEGER,
    vesting_counter SMALLINT,
    denom TEXT
);


CREATE TABLE accumulator_params
(
    one_row_id BOOLEAN NOT NULL DEFAULT TRUE PRIMARY KEY,
    params     JSONB   NOT NULL,
    CHECK (one_row_id)
);
-- +migrate Down
DROP TABLE admins_vesting;
DROP TABLE accumulator_params;
