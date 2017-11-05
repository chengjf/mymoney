CREATE TABLE t_record (
  `id`          INT UNSIGNED AUTO_INCREMENT NOT NULL
  COMMENT '自增主键',
  `type`        INT                         NOT NULL
  COMMENT '借贷类型，1-借，2-贷',
  `account_id`  INT UNSIGNED                NOT NULL
  COMMENT '借贷账户',
  `entry_id`    INT UNSIGNED                NOT NULL
  COMMENT '会计科目',
  `amount`      DOUBLE                      NOT NULL DEFAULT 0.00
  COMMENT '金额',
  `datetime`    DATETIME                    NOT NULL DEFAULT CURRENT_TIMESTAMP
  COMMENT '日期',
  `counter`     VARCHAR(255)
  COMMENT '借贷对象',
  `create_time` DATETIME                    NOT NULL DEFAULT CURRENT_TIMESTAMP
  COMMENT '记录创建时间',
  `update_time` DATETIME                    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
  COMMENT '记录更新时间',
  PRIMARY KEY (id)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='借贷记录表';
