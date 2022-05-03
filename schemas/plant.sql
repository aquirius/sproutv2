
CREATE TABLE `plant` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `uuid` char(36) CHARACTER SET ascii COLLATE ascii_general_ci DEFAULT NULL,
  `basket_id` int unsigned DEFAULT NULL,
  `name` varchar(100) NOT NULL,
  `description` varchar(500) DEFAULT NULL,
  `preview_image_uuid` char(36) CHARACTER SET ascii COLLATE ascii_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `plant_uuid_IDX` (`uuid`) USING BTREE,
  KEY `plant_fk` (`basket_id`),
  CONSTRAINT `plant_fk` FOREIGN KEY (`basket_id`) REFERENCES `basket` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci