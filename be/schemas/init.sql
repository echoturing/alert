CREATE TABLE IF NOT EXISTS alert
(
    `id`          BIGINT              NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT 'primary key',
    `name`        VARCHAR(100)        NOT NULL DEFAULT '' COMMENT 'alert name',
    `channels`    MEDIUMTEXT          NOT NULL,
    `rule`        LONGTEXT            NOT NULL,
    `alertStatus` tinyint(4) unsigned NOT NULL DEFAULT '1',
    `status`      tinyint(4) unsigned NOT NULL DEFAULT '1',
    `createdAt`   TIMESTAMP           NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updatedAt`   TIMESTAMP           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='alert';


CREATE TABLE IF NOT EXISTS datasource
(
    `id`        BIGINT              NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT 'primary key',
    `name`      VARCHAR(100)        NOT NULL DEFAULT '' COMMENT 'alert name',
    `type`      tinyint(4) unsigned NOT NULL DEFAULT '1',
    `detail`    LONGTEXT            NOT NULL,
    `createdAt` TIMESTAMP           NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updatedAt` TIMESTAMP           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP

) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='datasource';


CREATE TABLE IF NOT EXISTS channel
(
    `id`        BIGINT              NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT 'primary key',
    `name`      VARCHAR(100)        NOT NULL DEFAULT '' COMMENT 'channel name',
    `type`      tinyint(4) unsigned NOT NULL DEFAULT '1',
    `detail`    LONGTEXT            NOT NULL,
    `createdAt` TIMESTAMP           NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updatedAt` TIMESTAMP           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_bin COMMENT ='channel';