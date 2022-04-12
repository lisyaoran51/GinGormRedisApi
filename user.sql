CREATE DATABASE go_test_models;

USE go_test_models;

CREATE TABLE `users` (
	`userId` int NOT NULL AUTO_INCREMENT,
	`account` varchar(255),
	`password` varchar(255),
	`createdAt` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY(`userId`)
);