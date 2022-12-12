CREATE TABLE cola_types (
	id SERIAL NOT NULL,
	manufacturer_id INTEGER NOT NULL,
	name VARCHAR(255) NOT NULL,
	amount INTEGER NOT NULL,
	package_id INTEGER NOT NULL,
	calories INTEGER NOT NULL,
	PRIMARY KEY (id)
);