-- +goose Up
-- +goose StatementBegin
CREATE TABLE category (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT UNIQUE NOT NULL,
    estimated_months INTEGER NOT NULL,
    rate_indicated REAL,
    rules TEXT,
    notes TEXT
);

CREATE TABLE sub_category (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    rules TEXT,
    notes TEXT,
    category_id INTEGER NOT NULL,
    FOREIGN KEY (category_id) REFERENCES Category(id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE sub_category;

DROP TABLE category;
-- +goose StatementEnd
