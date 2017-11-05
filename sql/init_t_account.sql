insert into t_account (name, entry_id, balance) (select '现金账户1',id,0 from t_entry where name = '现金' and level = 2);
insert into t_account (name, entry_id, balance) (select '现金账户2',id,0 from t_entry where name = '现金' and level = 2);
insert into t_account (name, entry_id, balance) (select '虚拟支出账户',id,0 from t_entry where name = '收入类' and level = 1);
insert into t_account (name, entry_id, balance) (select '虚拟收入账户',id,0 from t_entry where name = '支出类' and level = 1);