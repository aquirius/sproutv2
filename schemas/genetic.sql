
CREATE TABLE `genetic` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `uuid` char(36) CHARACTER SET ascii COLLATE ascii_general_ci NOT NULL,
  `user_id` int unsigned NOT NULL,
  `name` varchar(100) NOT NULL,
  `description` varchar(500) NOT NULL,
  `customer_uuid` char(36) DEFAULT NULL,
  `created_ts` int unsigned NOT NULL,
  `updated_ts` int unsigned DEFAULT NULL,
  `preview_image_uuid` char(36) CHARACTER SET ascii COLLATE ascii_general_ci DEFAULT NULL,
  `station` int unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `genetic_uuid_IDX` (`uuid`) USING BTREE,
  KEY `genetic_fk` (`user_id`),
  CONSTRAINT `genetic_fk` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci
