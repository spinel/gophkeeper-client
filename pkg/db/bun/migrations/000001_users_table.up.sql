CREATE TABLE IF NOT EXISTS users (
	id serial PRIMARY KEY not null,
	email character varying unique,
	password text,
	created_at timestamp,
	updated_at timestamp,
	deleted_at timestamp
);
