-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS activities (
	id serial PRIMARY KEY,
	class_name text NOT NULL,
	title text NOT NULL,
	timestamp timestamptz NOT NULL DEFAULT NOW
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE activities;
-- +goose StatementEnd
