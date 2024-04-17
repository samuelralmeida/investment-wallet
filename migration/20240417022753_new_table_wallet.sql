-- +goose Up
-- +goose StatementBegin
CREATE TABLE wallets (
  id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  name TEXT UNIQUE NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE wallets;
-- +goose StatementEnd
