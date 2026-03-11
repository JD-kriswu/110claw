-- =============================================================
-- 110claw database schema
-- Engine: MariaDB 10.x+   Charset: utf8mb4_unicode_ci
-- Run via: ./initdb.sh   or manually: mysql < db/schema.sql
-- =============================================================

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- -------------------------------------------------------------
-- users  用户表
-- role: user | admin
-- status: 1=active  0=disabled
-- -------------------------------------------------------------
CREATE TABLE IF NOT EXISTS `users` (
  `id`            BIGINT UNSIGNED  NOT NULL AUTO_INCREMENT,
  `created_at`    DATETIME(3)      NULL,
  `updated_at`    DATETIME(3)      NULL,
  `deleted_at`    DATETIME(3)      NULL,
  `username`      VARCHAR(50)      NOT NULL,
  `phone`         VARCHAR(20)      NULL,
  `password_hash` VARCHAR(255)     NOT NULL,
  `avatar`        VARCHAR(500)     NULL,
  `role`          VARCHAR(20)      NOT NULL DEFAULT 'user',
  `status`        TINYINT          NOT NULL DEFAULT 1,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_users_username`   (`username`),
  UNIQUE KEY `uni_users_phone`      (`phone`),
  KEY           `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB
  DEFAULT CHARSET=utf8mb4
  COLLATE=utf8mb4_unicode_ci
  COMMENT='用户表';

-- -------------------------------------------------------------
-- news  资讯文章表
-- status: 1=published  0=draft
-- tags: JSON array stored as string, e.g. '["攻略","技巧"]'
-- -------------------------------------------------------------
CREATE TABLE IF NOT EXISTS `news` (
  `id`           BIGINT UNSIGNED  NOT NULL AUTO_INCREMENT,
  `created_at`   DATETIME(3)      NULL,
  `updated_at`   DATETIME(3)      NULL,
  `deleted_at`   DATETIME(3)      NULL,
  `title`        VARCHAR(200)     NOT NULL,
  `summary`      VARCHAR(500)     NULL,
  `content`      LONGTEXT         NULL,
  `cover_image`  VARCHAR(500)     NULL,
  `author_id`    BIGINT UNSIGNED  NULL,
  `tags`         VARCHAR(200)     NULL,
  `view_count`   BIGINT           NOT NULL DEFAULT 0,
  `status`       TINYINT          NOT NULL DEFAULT 1,
  `published_at` DATETIME(3)      NULL,
  PRIMARY KEY (`id`),
  KEY `idx_news_deleted_at`   (`deleted_at`),
  KEY `idx_news_status`       (`status`),
  KEY `idx_news_published_at` (`published_at`),
  KEY `idx_news_author_id`    (`author_id`)
) ENGINE=InnoDB
  DEFAULT CHARSET=utf8mb4
  COLLATE=utf8mb4_unicode_ci
  COMMENT='资讯文章表';

-- -------------------------------------------------------------
-- learn_steps  学习文章表（从飞书同步）
-- source_url: 飞书 wiki 页面 URL（唯一索引）
-- sort_order: 显示排序
-- status: 1=active  0=deleted
-- -------------------------------------------------------------
CREATE TABLE IF NOT EXISTS `learn_steps` (
  `id`          BIGINT UNSIGNED  NOT NULL AUTO_INCREMENT,
  `created_at`  DATETIME(3)      NULL,
  `updated_at`  DATETIME(3)      NULL,
  `deleted_at`  DATETIME(3)      NULL,
  `title`       VARCHAR(500)     NOT NULL,
  `description` VARCHAR(1000)    NULL,
  `content`     LONGTEXT         NULL,
  `source_url`  VARCHAR(500)     NULL,
  `sort_order`  BIGINT           NOT NULL DEFAULT 0,
  `status`      TINYINT          NOT NULL DEFAULT 1,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_learn_steps_source_url` (`source_url`),
  KEY        `idx_learn_steps_deleted_at` (`deleted_at`),
  KEY        `idx_learn_steps_sort_order` (`sort_order`)
) ENGINE=InnoDB
  DEFAULT CHARSET=utf8mb4
  COLLATE=utf8mb4_unicode_ci
  COMMENT='学习文章表';

-- -------------------------------------------------------------
-- skills  技能/插件目录表
-- rating: 0.0-5.0
-- tags: JSON array stored as string
-- status: 1=active  0=disabled
-- -------------------------------------------------------------
CREATE TABLE IF NOT EXISTS `skills` (
  `id`             BIGINT UNSIGNED  NOT NULL AUTO_INCREMENT,
  `created_at`     DATETIME(3)      NULL,
  `updated_at`     DATETIME(3)      NULL,
  `deleted_at`     DATETIME(3)      NULL,
  `name`           VARCHAR(100)     NOT NULL,
  `description`    VARCHAR(500)     NULL,
  `content`        LONGTEXT         NULL,
  `version`        VARCHAR(50)      NULL,
  `rating`         DOUBLE           NOT NULL DEFAULT 0,
  `download_count` BIGINT           NOT NULL DEFAULT 0,
  `author`         VARCHAR(100)     NULL,
  `tags`           VARCHAR(200)     NULL,
  `status`         TINYINT          NOT NULL DEFAULT 1,
  PRIMARY KEY (`id`),
  KEY `idx_skills_deleted_at`    (`deleted_at`),
  KEY `idx_skills_status`        (`status`),
  KEY `idx_skills_rating`        (`rating`),
  KEY `idx_skills_download_count`(`download_count`)
) ENGINE=InnoDB
  DEFAULT CHARSET=utf8mb4
  COLLATE=utf8mb4_unicode_ci
  COMMENT='技能/插件目录表';

SET FOREIGN_KEY_CHECKS = 1;
