ALTER TABLE
    `users` DROP INDEX `character_id`,
ADD
    INDEX `character_id` (`character_id`) USING BTREE;