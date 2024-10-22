-- +migrate Up
CREATE TABLE nft_events
(
    id                  SERIAL,
    event_type          TEXT    NOT NULL,
    nft_address         TEXT  REFERENCES nfts (address) NOT NULL,
    owner               TEXT    NOT NULL,
    new_owner           TEXT,
    validator           TEXT,
    amount              COIN[],

);
CREATE INDEX nft_events_id_index ON nft_events (id);

CREATE TABLE nfts
(
    address             TEXT    NOT NULL UNIQUE,
    owner               TEXT    NOT NULL,
    locked_amount       COIN[],
    available_amount    COIN[],
    delegations         delegation[]
);

CREATE TYPE delegation AS
(
    validator TEXT,
    amount TEXT,
    timestamp TIMESTAMP WITHOUT TIME ZONE
);

-- +migrate Down
DROP TABLE nft_events;
DROP TABLE nfts;
