CREATE TABLE users (
	id SERIAL NOT NULL,
	name VARCHAR(255) NOT NULL,
	email VARCHAR(255) NOT NULL,
	hashed_password VARCHAR(255) NOT NULL,
	role_id INTEGER NOT NULL,
	biography VARCHAR(255),
	user_icon VARCHAR(255),
	PRIMARY KEY (id)
);