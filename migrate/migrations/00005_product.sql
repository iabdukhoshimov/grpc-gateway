-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS products (
	id serial PRIMARY KEY NOT NULL,
	name varchar NOT NULL,
	price numeric NOT NULL,
	stock int NOT NULL,
	created_at timestamp WITH time zone NOT NULL DEFAULT (NOW()),
	updated_at timestamp WITH time zone,
	expired_at timestamp WITH time zone,
	deleted_at timestamp WITH time zone DEFAULT NULL
);


-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS products;

-- +goose StatementEnd