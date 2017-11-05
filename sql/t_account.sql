CREATE TABLE t_account (
  `id`          INT UNSIGNED AUTO_INCREMENT NOT NULL
  COMMENT '自增主键',
  `name`        VARCHAR(255)                NOT NULL
  COMMENT '账户名称',
  `entry_id`    INT UNSIGNED                NOT NULL
  COMMENT '会计科目，二级科目',
  `balance`     INT                         NOT NULL DEFAULT 0
  COMMENT '金额',
  `create_time` DATETIME                    NOT NULL DEFAULT CURRENT_TIMESTAMP
  COMMENT '记录创建时间',
  `update_time` DATETIME                    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
  COMMENT '记录更新时间',
  PRIMARY KEY (id),
  UNIQUE KEY (name)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='借贷账户表';
