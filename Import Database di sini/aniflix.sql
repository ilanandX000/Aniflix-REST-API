-- Adminer 4.8.1 MySQL 10.11.6-MariaDB-0+deb12u1 dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

CREATE DATABASE `aniflix` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci */;
USE `aniflix`;

DROP TABLE IF EXISTS `anime`;
CREATE TABLE `anime` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `synopsis` varchar(255) DEFAULT NULL,
  `thumbnail` varchar(255) DEFAULT NULL,
  `rating` float DEFAULT NULL,
  `status` int(11) DEFAULT NULL COMMENT '1 = ongoing, 2 = complete',
  `release_date` date DEFAULT NULL,
  `created_at` int(11) DEFAULT NULL,
  `updated_at` int(11) DEFAULT NULL,
  `created_by` int(11) DEFAULT NULL,
  `updated_by` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO `anime` (`id`, `title`, `synopsis`, `thumbnail`, `rating`, `status`, `release_date`, `created_at`, `updated_at`, `created_by`, `updated_by`) VALUES
(1,	'Movie Title',	'Movie synopsis goes here',	'http://example.com/thumbnail.jpg',	4.5,	1,	'2025-01-08',	1633046400,	1736309292,	1,	1),
(2,	'One Piece',	'Petualangan Luffy untuk menjadi Raja Bajak Laut.',	'http://example.com/onepiece.jpg',	8.9,	1,	'2002-10-03',	NULL,	NULL,	1,	1),
(4,	'Solo Leveling',	'Movie synopsis goes here',	'http://example.com/thumbnail.jpg',	9.5,	1,	'2024-01-20',	1633046400,	1633046400,	1,	1);

DROP TABLE IF EXISTS `anime_episode`;
CREATE TABLE `anime_episode` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `anime_id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `sequence` int(11) NOT NULL,
  `stream_url` varchar(255) DEFAULT NULL,
  `created_at` int(11) DEFAULT NULL,
  `updated_at` int(11) DEFAULT NULL,
  `created_by` int(11) DEFAULT NULL,
  `updated_by` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `anime_id` (`anime_id`),
  CONSTRAINT `anime_episode_ibfk_1` FOREIGN KEY (`anime_id`) REFERENCES `anime` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO `anime_episode` (`id`, `anime_id`, `name`, `sequence`, `stream_url`, `created_at`, `updated_at`, `created_by`, `updated_by`) VALUES
(1,	2,	'Naruto',	5,	NULL,	NULL,	NULL,	NULL,	NULL);

DROP TABLE IF EXISTS `anime_genre`;
CREATE TABLE `anime_genre` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `anime_id` int(11) NOT NULL,
  `genre_id` int(11) NOT NULL,
  `created_at` int(11) DEFAULT NULL,
  `updated_at` int(11) DEFAULT NULL,
  `created_by` int(11) DEFAULT NULL,
  `updated_by` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `anime_id` (`anime_id`),
  KEY `genre_id` (`genre_id`),
  CONSTRAINT `anime_genre_ibfk_1` FOREIGN KEY (`anime_id`) REFERENCES `anime` (`id`),
  CONSTRAINT `anime_genre_ibfk_2` FOREIGN KEY (`genre_id`) REFERENCES `genre` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO `anime_genre` (`id`, `anime_id`, `genre_id`, `created_at`, `updated_at`, `created_by`, `updated_by`) VALUES
(4,	1,	3,	1633046400,	1633046400,	1,	1);

DROP TABLE IF EXISTS `genre`;
CREATE TABLE `genre` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `created_at` int(11) DEFAULT NULL,
  `updated_at` int(11) DEFAULT NULL,
  `created_by` int(11) DEFAULT NULL,
  `updated_by` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO `genre` (`id`, `name`, `created_at`, `updated_at`, `created_by`, `updated_by`) VALUES
(1,	'Fantasy',	1633046400,	1736310421,	1,	1),
(3,	'Isekai',	1633046400,	1633046400,	1,	1);

-- 2025-01-08 08:12:35