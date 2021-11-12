CREATE TABLE `characters` (
    `id` INT UNSIGNED NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    `corporation_id` INT UNSIGNED NOT NULL,
    `alliance_id` INT UNSIGNED NULL DEFAULT NULL,
    `faction_id` INT UNSIGNED NULL DEFAULT NULL,
    `security_status` FLOAT NULL DEFAULT NULL,
    `gender` ENUM('male', 'female') NOT NULL,
    `birthday` DATETIME NOT NULL,
    `title` VARCHAR(255) NULL DEFAULT NULL,
    `bloodline_id` INT UNSIGNED NOT NULL,
    `race_id` INT UNSIGNED NOT NULL,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    PRIMARY KEY (`id`) USING BTREE,
    INDEX `characters_corporation_id_idx` (`corporation_id`) USING BTREE,
    INDEX `characters_alliance_id_idx` (`alliance_id`) USING BTREE,
    INDEX `characters_faction_id_idx` (`faction_id`) USING BTREE,
    INDEX `characters_bloodline_id_idx` (`bloodline_id`) USING BTREE,
    INDEX `characters_race_id_idx` (`race_id`) USING BTREE
) COLLATE = 'utf8mb4_unicode_ci' ENGINE = InnoDB;