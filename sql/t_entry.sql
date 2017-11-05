CREATE TABLE t_entry (
  `id`          INT UNSIGNED AUTO_INCREMENT NOT NULL
  COMMENT '自增主键',
  `name`        VARCHAR(255)                NOT NULL
  COMMENT '名称',
  `level`       INT UNSIGNED                NOT NULL DEFAULT 0
  COMMENT '级别',
  `parent_lvl`  INT UNSIGNED
  COMMENT '父级别',
  `create_time` DATETIME                    NOT NULL DEFAULT CURRENT_TIMESTAMP
  COMMENT '记录创建时间',
  `update_time` DATETIME                    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
  COMMENT '记录更新时间',
  PRIMARY KEY (id),
  UNIQUE KEY (name, level, parent_lvl)
)
  ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COMMENT ='会计科目表';
