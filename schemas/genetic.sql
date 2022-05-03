
CREATE TABLE `genetic` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `uuid` char(36) CHARACTER SET ascii COLLATE ascii_general_ci NOT NULL,
  `plant_id` int unsigned NOT NULL,
  `ph` varchar(100) NOT NULL,
  `ec` varchar(500) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `genetic_uuid_IDX` (`uuid`) USING BTREE,
  KEY `genetic_fk` (`plant_id`),
  CONSTRAINT `genetic_fk` FOREIGN KEY (`plant_id`) REFERENCES `plant` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci
