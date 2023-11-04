CREATE TABLE urls(
	id serial,
	originalUrl TEXT NOT NULL,
	hash VARCHAR(20) NOT NULL,
	created_at BIGINT,
	owner BIGINT NULL,
	PRIMARY KEY(id),
	FOREIGN KEY(owner) REFERENCES users(id)
);

CREATE TABLE users(
	id serial,
	name VARCHAR(200),
	PRIMARY KEY(id)
);