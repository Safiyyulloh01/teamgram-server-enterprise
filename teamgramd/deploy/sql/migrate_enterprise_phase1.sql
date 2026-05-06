-- Teamgram Enterprise Features - Phase 1 Database Migrations
-- Run: mysql -uroot teamgram < migrate_enterprise_phase1.sql

-- Channels & Megagroups
CREATE TABLE IF NOT EXISTS `channels` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `channel_id` BIGINT NOT NULL DEFAULT 0,
    `access_hash` BIGINT NOT NULL DEFAULT 0,
    `title` VARCHAR(255) NOT NULL DEFAULT '',
    `about` TEXT,
    `photo` TEXT,
    `creator_user_id` BIGINT NOT NULL DEFAULT 0,
    `participants_count` INT NOT NULL DEFAULT 0,
    `admins_count` INT NOT NULL DEFAULT 0,
    `kicked_count` INT NOT NULL DEFAULT 0,
    `banned_count` INT NOT NULL DEFAULT 0,
    `date` INT NOT NULL DEFAULT 0,
    `version` INT NOT NULL DEFAULT 0,
    `username` VARCHAR(255) NOT NULL DEFAULT '',
    `signatures` TINYINT(1) NOT NULL DEFAULT 0,
    `signature_profiles` TINYINT(1) NOT NULL DEFAULT 0,
    `slow_mode_seconds` INT NOT NULL DEFAULT 0,
    `linked_chat_id` BIGINT NOT NULL DEFAULT 0,
    `location` TEXT,
    `has_link` TINYINT(1) NOT NULL DEFAULT 0,
    `democracy` TINYINT(1) NOT NULL DEFAULT 0,
    `migrated_from_chat_id` BIGINT NOT NULL DEFAULT 0,
    `migrated_from_max_id` INT NOT NULL DEFAULT 0,
    `deleted` TINYINT(1) NOT NULL DEFAULT 0,
    `default_banned_rights` INT NOT NULL DEFAULT 0,
    `participants_type` INT NOT NULL DEFAULT 0,
    `pinned_msg_id` INT NOT NULL DEFAULT 0,
    `ttl_period` INT NOT NULL DEFAULT 0,
    `theme_emoticon` VARCHAR(255) NOT NULL DEFAULT '',
    `available_reactions_type` INT NOT NULL DEFAULT 0,
    `available_reactions` TEXT,
    `sticker_set_id` BIGINT NOT NULL DEFAULT 0,
    `can_set_sticker_set` TINYINT(1) NOT NULL DEFAULT 0,
    `send_message_level` INT NOT NULL DEFAULT 0,
    `converted_to_gigagroup` TINYINT(1) NOT NULL DEFAULT 0,
    `no_forwards` TINYINT(1) NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`),
    INDEX `idx_channel_id` (`channel_id`),
    INDEX `idx_username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `channel_participants` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `channel_id` BIGINT NOT NULL DEFAULT 0,
    `user_id` BIGINT NOT NULL DEFAULT 0,
    `participant_type` INT NOT NULL DEFAULT 0,
    `promoted_by` BIGINT NOT NULL DEFAULT 0,
    `rank` VARCHAR(255) NOT NULL DEFAULT '',
    `date` INT NOT NULL DEFAULT 0,
    `inviter_id` BIGINT NOT NULL DEFAULT 0,
    `kicked_by` BIGINT NOT NULL DEFAULT 0,
    `banned_rights` INT NOT NULL DEFAULT 0,
    `until_date` INT NOT NULL DEFAULT 0,
    `is_migrated` TINYINT(1) NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`),
    INDEX `idx_channel_id` (`channel_id`),
    INDEX `idx_user_id` (`user_id`),
    INDEX `idx_channel_user` (`channel_id`, `user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Bots
CREATE TABLE IF NOT EXISTS `bot_data` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `bot_id` BIGINT NOT NULL DEFAULT 0,
    `bot_token` VARCHAR(255) NOT NULL DEFAULT '',
    `bot_type` INT NOT NULL DEFAULT 0,
    `creator_user_id` BIGINT NOT NULL DEFAULT 0,
    `description` TEXT,
    `inline_placeholder` VARCHAR(255) NOT NULL DEFAULT '',
    `inline_geo` TINYINT(1) NOT NULL DEFAULT 0,
    `inline_js` TINYINT(1) NOT NULL DEFAULT 0,
    `can_see_history` TINYINT(1) NOT NULL DEFAULT 0,
    `can_join_groups` TINYINT(1) NOT NULL DEFAULT 1,
    `can_read_all_messages` TINYINT(1) NOT NULL DEFAULT 0,
    `privacy_settings` INT NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`),
    INDEX `idx_bot_id` (`bot_id`),
    INDEX `idx_bot_token` (`bot_token`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Bot commands
CREATE TABLE IF NOT EXISTS `bot_commands` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `bot_id` BIGINT NOT NULL DEFAULT 0,
    `user_id` BIGINT NOT NULL DEFAULT 0,
    `scope` INT NOT NULL DEFAULT 0,
    `command` VARCHAR(64) NOT NULL DEFAULT '',
    `description` VARCHAR(255) NOT NULL DEFAULT '',
    PRIMARY KEY (`id`),
    INDEX `idx_bot_user` (`bot_id`, `user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Reactions
CREATE TABLE IF NOT EXISTS `message_reactions` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT NOT NULL DEFAULT 0,
    `peer_type` INT NOT NULL DEFAULT 0,
    `peer_id` BIGINT NOT NULL DEFAULT 0,
    `message_id` INT NOT NULL DEFAULT 0,
    `reaction` VARCHAR(64) NOT NULL DEFAULT '',
    `reaction_date` INT NOT NULL DEFAULT 0,
    `is_big` TINYINT(1) NOT NULL DEFAULT 0,
    `deleted` TINYINT(1) NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`),
    INDEX `idx_message` (`peer_type`, `peer_id`, `message_id`),
    INDEX `idx_user_message` (`user_id`, `message_id`),
    UNIQUE KEY `uk_user_message_reaction` (`user_id`, `peer_type`, `peer_id`, `message_id`, `reaction`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Sticker sets
CREATE TABLE IF NOT EXISTS `sticker_sets` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `sticker_set_id` BIGINT NOT NULL DEFAULT 0,
    `access_hash` BIGINT NOT NULL DEFAULT 0,
    `title` VARCHAR(255) NOT NULL DEFAULT '',
    `short_name` VARCHAR(255) NOT NULL DEFAULT '',
    `count` INT NOT NULL DEFAULT 0,
    `hash` INT NOT NULL DEFAULT 0,
    `date` INT NOT NULL DEFAULT 0,
    `is_archived` TINYINT(1) NOT NULL DEFAULT 0,
    `is_official` TINYINT(1) NOT NULL DEFAULT 0,
    `is_masks` TINYINT(1) NOT NULL DEFAULT 0,
    `is_emojis` TINYINT(1) NOT NULL DEFAULT 0,
    `is_animated` TINYINT(1) NOT NULL DEFAULT 0,
    `is_videos` TINYINT(1) NOT NULL DEFAULT 0,
    `thumbnail` TEXT,
    `thumbnail_dc_id` INT NOT NULL DEFAULT 0,
    `installed_count` INT NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`),
    INDEX `idx_sticker_set_id` (`sticker_set_id`),
    INDEX `idx_short_name` (`short_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `sticker_pack` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `sticker_set_id` BIGINT NOT NULL DEFAULT 0,
    `emoticon` VARCHAR(64) NOT NULL DEFAULT '',
    `document_id` BIGINT NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`),
    INDEX `idx_sticker_set_id` (`sticker_set_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS `user_sticker_installs` (
    `id` BIGINT NOT NULL AUTO_INCREMENT,
    `user_id` BIGINT NOT NULL DEFAULT 0,
    `sticker_set_id` BIGINT NOT NULL DEFAULT 0,
    `order_num` INT NOT NULL DEFAULT 0,
    `is_archived` TINYINT(1) NOT NULL DEFAULT 0,
    `date` INT NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`),
    INDEX `idx_user_id` (`user_id`),
    INDEX `idx_user_set` (`user_id`, `sticker_set_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
