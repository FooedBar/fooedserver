CREATE DATABASE IF NOT EXISTS finder;

USE finder;

DROP TABLE IF EXISTS `selection`, `menu_item`, `restaurant`, `session`;

CREATE TABLE IF NOT EXISTS `session` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `current_lat` float,
    `current_long` float,
    `created_at` datetime,
    `updated_at` datetime,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `restaurant` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `lat` float,
    `long` float,
    `name` varchar(255),
    `description` varchar(255),
    `style` varchar(200),
    `image_url` varchar(255),
    `image_width` int,
    `image_height` int,
    `created_at` datetime,
    `updated_at` datetime,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `menu_item` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `restaurant_id` bigint(20) NOT NULL,
    `image_url` varchar(255),
    `image_height` int,
    `image_width` int,
    `name` varchar(255),
    `description` varchar(255),
    `flavor` varchar(100),
    `heavy` float,
    `style_one` varchar(200),
    `style_two` varchar(200),
    `style_three` varchar(200),
    `created_at` datetime,
    `updated_at` datetime,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`restaurant_id`) REFERENCES restaurant(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `selection` (
	`id` bigint(20) NOT NULL AUTO_INCREMENT,
    `session_id` bigint(20) NOT NULL,
    `menu_item_id` bigint(20) NOT NULL,
    `like` tinyint(1),
    PRIMARY KEY (`id`),
    FOREIGN KEY (`session_id`) REFERENCES session(id),
    FOREIGN KEY (`menu_item_id`) REFERENCES menu_item(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
