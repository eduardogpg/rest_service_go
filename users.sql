Create database test;
Use test;

CREATE TABLE users
(
	Id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
	created_at  datetime,
	update_at datetime,
	user_name varchar(64)
);