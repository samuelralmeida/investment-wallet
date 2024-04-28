-- +goose Up
-- +goose StatementBegin
CREATE TABLE funds (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    cnpj TEXT NOT NULL,
    bank TEXT NOT NULL,
    min_value REAL NOT NULL,
    notes TEXT,
    benchmark TEXT,
    subcategory_id INTEGER NOT NULL,
    FOREIGN KEY (subcategory_id) REFERENCES sub_category(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE funds;
-- +goose StatementEnd
