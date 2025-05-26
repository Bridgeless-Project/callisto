-- +migrate Up
ALTER TABLE proposal
ALTER COLUMN content TYPE TEXT USING content::text,
    ADD COLUMN metadata TEXT;

-- +migrate Down
ALTER TABLE proposal
ALTER COLUMN content TYPE JSONB USING content::jsonb,
    DROP COLUMN metadata;