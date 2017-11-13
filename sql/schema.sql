CREATE DATABASE IF NOT EXISTS teleport CHARACTER SET utf8 COLLATE utf8_general_ci;

use teleport;

CREATE TABLE `users` (
  `login` CHAR(36) NOT NULL,
  `pass` CHAR(36) NOT NULL,
  `payload`    JSON COMMENT 'Данные',
  `created_at` DATETIME NOT NULL COMMENT 'Дата создания',
  `expired_at` DATETIME NOT NULL COMMENT 'Дата окончания',
  `active` TINYINT(1) DEFAULT 0,
  PRIMARY KEY (`login`)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8
  COMMENT 'teleport data';
