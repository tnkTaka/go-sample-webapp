-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE `tasks` (
  `id` INT(4) NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT 'ID',
  `text` VARCHAR(20) NOT NULL COMMENT '内容',
  `create_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日',
  `update_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日'
);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `tasks`;