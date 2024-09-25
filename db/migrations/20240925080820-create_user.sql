
-- +migrate Up
CREATE TABLE IF NOT EXISTS user {
}
-- +migrate Down
DROP TABLE IF EXISTS user;
