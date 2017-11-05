## 获取所有的资产类账户
SELECT a.*
FROM t_account a, t_entry b, t_entry c
WHERE a.entry_id = b.id AND b.parent_lvl = c.id AND c.name = '资产类';

## 获取所有的顶级科目
SELECT a.*
FROM t_entry a
WHERE level = 1;

## 获取所有的支出科目
SELECT c.*
FROM t_entry a, t_entry b, t_entry c
WHERE a.name = '支出类'
      AND b.parent_lvl = a.id
      AND c.parent_lvl = b.id;