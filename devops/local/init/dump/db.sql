CREATE TABLE `db`.users (
	id TINYINT UNSIGNED auto_increment NOT NULL,
	name varchar(100) NOT NULL,
	age varchar(100) NOT NULL,
	email varchar(100) NOT NULL,
	password varchar(100) NOT NULL,
	address varchar(100) NOT NULL,
	PRIMARY KEY (id)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4;