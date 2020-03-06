SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article`
(
    `id`            INT(10) UNSIGNED    NOT NULL AUTO_INCREMENT,
    `category_id`   INT(10) UNSIGNED    NOT NULL,
    `summary`       VARCHAR(255)        NOT NULL DEFAULT '' COMMENT '摘要',
    `content`       TEXT                NOT NULL,
    `title`         VARCHAR(100)        NOT NULL DEFAULT '',
    `view_count`    INT(11)                      DEFAULT '0',
    `comment_count` INT(11) UNSIGNED             DEFAULT NULL COMMENT '0',
    `username`      VARCHAR(80)                  DEFAULT '',
    `status`        TINYINT(1) UNSIGNED NOT NULL DEFAULT '1',
    # `create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, # required 5.6
    `create_time`   DATETIME            NOT NULL,
    `update_time`   TIMESTAMP           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = INNODB
  AUTO_INCREMENT = 3
  DEFAULT CHARSET = utf8mb4;

-- ----------------------------
-- Records of article
-- ----------------------------
INSERT INTO `article`
VALUES ('1', '1', 'sum1', '测试1测试1测试1', '测试1', '0', '0', 'haima', '1', '2019-11-10 09:48:40', '2019-11-10 09:48:40');
INSERT INTO `article`
VALUES ('2', '1', 'sum2', '测试2测试2测试2', '测试2', '0', '0', 'haima', '1', '2019-11-10 10:38:53', '2019-11-10 10:38:53');

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category`
(
    `id`            BIGINT(20) UNSIGNED NOT NULL,
    `category_name` VARCHAR(255)        NOT NULL DEFAULT '',
    `category_no`   INT(10) UNSIGNED             DEFAULT NULL,
    # `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, # required 5.6
    `create_time`   DATETIME            NULL,
    `update_time`   TIMESTAMP           NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = INNODB
  DEFAULT CHARSET = utf8mb4;

-- ----------------------------
-- Records of category
-- ----------------------------
INSERT INTO `category`
VALUES ('1', '分类一', '10001', '2019-11-10 09:12:16', '2019-11-10 09:12:16');
INSERT INTO `category`
VALUES ('2', '分类二', '10002', '2019-11-10 09:24:52', '2019-11-10 09:24:52');

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment`
(
    `id`          BIGINT(20)          NOT NULL,
    `content`     VARCHAR(255)        NOT NULL,
    `username`    VARCHAR(60)         NOT NULL DEFAULT '',
    `status`      TINYINT(1) UNSIGNED NOT NULL DEFAULT '1',
    `article_id`  BIGINT(20)          NOT NULL,
    `create_time` TIMESTAMP           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = INNODB
  DEFAULT CHARSET = utf8mb4;

-- ----------------------------
-- Records of comment
-- ----------------------------

-- ----------------------------
-- Table structure for leave
-- ----------------------------
DROP TABLE IF EXISTS `leave`;
CREATE TABLE `leave`
(
    `id`          BIGINT(20)   NOT NULL,
    `username`    VARCHAR(60)  NOT NULL,
    `email`       VARCHAR(160) NOT NULL DEFAULT '',
    `content`     VARCHAR(255)          DEFAULT '',
    # `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, # required 5.6
    `create_time` DATETIME     NULL,
    `update_time` TIMESTAMP    NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = INNODB
  DEFAULT CHARSET = utf8mb4;

-- ----------------------------
-- Records of leave
-- ----------------------------
