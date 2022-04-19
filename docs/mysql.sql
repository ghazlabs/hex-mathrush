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

INSERT INTO questions (problem, correct_index) VALUES
('1 + 1', 2),
('1 + 2', 3),
('2 + 1', 3),
('3 - 2', 1),
('2 - 1', 1),
('2 + 2 - 1', 3),
('1 + 1 + 1', 3),
('3 - 1', 2),
('2 + 1 - 2', 1),
('1 + 1 - 1', 1);

INSERT INTO choices (question_id, choice) VALUES
(1, '1'),
(1, '2'),
(1, '3'),
(2, '1'),
(2, '2'),
(2, '3'),
(3, '1'),
(3, '2'),
(3, '3'),
(4, '1'),
(4, '2'),
(4, '3'),
(5, '1'),
(5, '2'),
(5, '3'),
(6, '1'),
(6, '2'),
(6, '3'),
(7, '1'),
(7, '2'),
(7, '3'),
(8, '1'),
(8, '2'),
(8, '3'),
(9, '1'),
(9, '2'),
(9, '3'),
(10, '1'),
(10, '2'),
(10, '3'),;
