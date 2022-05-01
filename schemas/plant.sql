

CREATE TABLE `plant` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `uuid` char(36) CHARACTER SET ascii COLLATE ascii_general_ci NOT NULL,
  `name` varchar(100) NOT NULL,
  `description` varchar(500) NOT NULL,
  `created_ts` int unsigned NOT NULL,
  `updated_ts` int unsigned DEFAULT NULL,
  `preview_image_uuid` char(36) CHARACTER SET ascii COLLATE ascii_general_ci DEFAULT NULL,
  `station` int unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `plant_uuid_IDX` (`uuid`) USING BTREE,
  KEY `plant_fk` (`user_id`),
  CONSTRAINT `plant_fk` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci
