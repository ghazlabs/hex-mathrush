CREATE DATABASE hex_math;
USE hex_math;

CREATE TABLE `questions` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `problem` text NOT NULL,
  `correct_index` int NOT NULL
);

CREATE TABLE `choices` (
  `question_id` int,
  `choice` varchar(255),
  PRIMARY KEY (`question_id`, `choice`)
);

CREATE TABLE `games` (
  `id` varchar(36)  PRIMARY KEY,
  `player_name` varchar(36) NOT NULL,
  `scenario` varchar(20) NOT NULL,
  `score` int NOT NULL,
  `count_correct` int NOT NULL,
  `question_id` int NOT NULL,
  `question_timeout` int NOT NULL
);

ALTER TABLE `choices` ADD FOREIGN KEY (`question_id`) REFERENCES `questions` (`id`);

ALTER TABLE `games` ADD FOREIGN KEY (`question_id`) REFERENCES `questions` (`id`);

