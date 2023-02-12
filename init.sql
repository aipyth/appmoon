-- CREATE ROLE appmoon LOGIN PASSWORD :POSTGRES_PASSWORD;
-- GRANT CONNECT ON DATABASE postgres TO appmoon;
-- GRANT USAGE, SELECT ON SCHEMA public TO appmoon;
-- ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT ON TABLES TO appmoon;

CREATE TABLE IF NOT EXISTS activities (
	id serial PRIMARY KEY,
	class_name text NOT NULL,
	title text NOT NULL,
	timestamp timestamptz NOT NULL
);
