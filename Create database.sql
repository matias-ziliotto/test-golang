CREATE TABLE IF NOT EXISTS document_types(
	id INT NOT NULL AUTO_INCREMENT,
	name VARCHAR (255) NOT NULL,
	code VARCHAR (255) NOT NULL,

	PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS roles(
	id INT NOT NULL AUTO_INCREMENT,
	name VARCHAR (255) NOT NULL,
	code VARCHAR (255) NOT NULL,

	PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS users(
	id INT NOT NULL AUTO_INCREMENT,
	first_name VARCHAR (255) NOT NULL,
	last_name VARCHAR (255) NOT NULL,
    document_type_id int NOT NULL,
    document_number int NOT NULL,

	PRIMARY KEY(id),
	FOREIGN KEY (document_type_id) REFERENCES document_types(id)
);

CREATE TABLE IF NOT EXISTS roles_users(
	role_id INT NOT NULL,
	user_id INT NOT NULL,

	PRIMARY KEY(role_id, user_id),
	FOREIGN KEY (role_id) REFERENCES roles(id),
	FOREIGN KEY (user_id) REFERENCES users(id)
);
