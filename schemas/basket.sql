
CREATE TABLE `basket` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `uuid` char(36) CHARACTER SET ascii COLLATE ascii_general_ci DEFAULT NULL,
  `user_id` int unsigned NOT NULL,
  `name` varchar(100) NOT NULL,
  `description` varchar(500) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `basket_uuid_IDX` (`uuid`) USING BTREE,
  KEY `basket_fk` (`user_id`),
  CONSTRAINT `basket_fk` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci