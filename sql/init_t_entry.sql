insert into t_entry(name, level, parent_lvl) value ('资产类',1,null);
insert into t_entry(name, level, parent_lvl) (select '现金',2,id from t_entry where name = '资产类' and level = 1);
insert into t_entry(name, level, parent_lvl) (select '银行存款',2,id from t_entry where name = '资产类' and level = 1);
insert into t_entry(name, level, parent_lvl) (select '应收利息',2,id from t_entry where name = '资产类' and level = 1);
insert into t_entry(name, level, parent_lvl) (select '应收账款',2,id from t_entry where name = '资产类' and level = 1);
insert into t_entry(name, level, parent_lvl) (select '其他',2,id from t_entry where name = '资产类' and level = 1);

insert into t_entry(name, level, parent_lvl) value ('负债类',1,null);
insert into t_entry(name, level, parent_lvl) (select '应付账款',2,id from t_entry where name = '负债类' and level = 1);
insert into t_entry(name, level, parent_lvl) (select '应交税金',2,id from t_entry where name = '负债类' and level = 1);
insert into t_entry(name, level, parent_lvl) (select '其他',2,id from t_entry where name = '负债类' and level = 1);

insert into t_entry(name, level, parent_lvl) value ('收入类',1,null);
insert into t_entry(name, level, parent_lvl) (select '工资收入',2,id from t_entry where name = '收入类' and level = 1);
insert into t_entry(name, level, parent_lvl) (select '奖金',2,id from t_entry where name = '收入类' and level = 1);
insert into t_entry(name, level, parent_lvl) (select '加班费',2,id from t_entry where name = '收入类' and level = 1);
insert into t_entry(name, level, parent_lvl) (select '投资收益',2,id from t_entry where name = '收入类' and level = 1);
insert into t_entry(name, level, parent_lvl) (select '其他',2,id from t_entry where name = '收入类' and level = 1);
 

insert into t_entry(name, level, parent_lvl) value ('支出类',1,null);
insert into t_entry(name, level, parent_lvl) (select '食品酒水',2,id from t_entry where name = '支出类' and level = 1);
insert into t_entry(name, level, parent_lvl) (select '早餐',3,id from t_entry where name = '食品酒水' and level = 2);
insert into t_entry(name, level, parent_lvl) (select '午餐',3,id from t_entry where name = '食品酒水' and level = 2);
insert into t_entry(name, level, parent_lvl) (select '晚餐',3,id from t_entry where name = '食品酒水' and level = 2);
insert into t_entry(name, level, parent_lvl) (select '下午茶',3,id from t_entry where name = '食品酒水' and level = 2);
insert into t_entry(name, level, parent_lvl) (select '水果',3,id from t_entry where name = '食品酒水' and level = 2);
insert into t_entry(name, level, parent_lvl) (select '饮料酒水',3,id from t_entry where name = '食品酒水' and level = 2);
insert into t_entry(name, level, parent_lvl) (select '食品零食',3,id from t_entry where name = '食品酒水' and level = 2);
insert into t_entry(name, level, parent_lvl) (select '蔬菜原料',3,id from t_entry where name = '食品酒水' and level = 2);
insert into t_entry(name, level, parent_lvl) (select '油盐酱醋',3,id from t_entry where name = '食品酒水' and level = 2);
insert into t_entry(name, level, parent_lvl) (select '其他',3,id from t_entry where name = '食品酒水' and level = 2);

insert into t_entry(name, level, parent_lvl) (select '账单',2,id from t_entry where name = '支出类' and level = 1);
insert into t_entry(name, level, parent_lvl) (select '房租',3,id from t_entry where name = '账单' and level = 2);
insert into t_entry(name, level, parent_lvl) (select '水费',3,id from t_entry where name = '账单' and level = 2);
insert into t_entry(name, level, parent_lvl) (select '电费',3,id from t_entry where name = '账单' and level = 2);
insert into t_entry(name, level, parent_lvl) (select '燃气费',3,id from t_entry where name = '账单' and level = 2);
insert into t_entry(name, level, parent_lvl) (select '管理费',3,id from t_entry where name = '账单' and level = 2);
insert into t_entry(name, level, parent_lvl) (select '网络宽带',3,id from t_entry where name = '账单' and level = 2);
insert into t_entry(name, level, parent_lvl) (select '话费',3,id from t_entry where name = '账单' and level = 2);
insert into t_entry(name, level, parent_lvl) (select '电视费',3,id from t_entry where name = '账单' and level = 2);
insert into t_entry(name, level, parent_lvl) (select 'VPS',3,id from t_entry where name = '账单' and level = 2);
insert into t_entry(name, level, parent_lvl) (select '域名',3,id from t_entry where name = '账单' and level = 2);
insert into t_entry(name, level, parent_lvl) (select '会员',3,id from t_entry where name = '账单' and level = 2);
insert into t_entry(name, level, parent_lvl) (select '订阅服务',3,id from t_entry where name = '账单' and level = 2);
insert into t_entry(name, level, parent_lvl) (select '其他',3,id from t_entry where name = '账单' and level = 2);

insert into t_entry(name, level, parent_lvl) (select '人情关系',2,id from t_entry where name = '支出类' and level = 1);
insert into t_entry(name, level, parent_lvl) (select '结婚礼金',3,id from t_entry where name = '人情关系' and level = 2);
insert into t_entry(name, level, parent_lvl) (select '红包',3,id from t_entry where name = '人情关系' and level = 2);
insert into t_entry(name, level, parent_lvl) (select '聚餐',3,id from t_entry where name = '人情关系' and level = 2);
insert into t_entry(name, level, parent_lvl) (select '礼物',3,id from t_entry where name = '人情关系' and level = 2);
insert into t_entry(name, level, parent_lvl) (select '孝敬',3,id from t_entry where name = '人情关系' and level = 2);
insert into t_entry(name, level, parent_lvl) (select '其他',3,id from t_entry where name = '人情关系' and level = 2);

insert into t_entry(name, level, parent_lvl) (select '行车交通',2,id from t_entry where name = '支出类' and level = 1);
insert into t_entry(name, level, parent_lvl) (select '公交地铁',3,id from t_entry where name = '行车交通' and level = 2);
insert into t_entry(name, level, parent_lvl) (select '出租车',3,id from t_entry where name = '行车交通' and level = 2);
insert into t_entry(name, level, parent_lvl) (select '客车',3,id from t_entry where name = '行车交通' and level = 2);
insert into t_entry(name, level, parent_lvl) (select '高铁火车',3,id from t_entry where name = '行车交通' and level = 2);
insert into t_entry(name, level, parent_lvl) (select '飞机',3,id from t_entry where name = '行车交通' and level = 2);
insert into t_entry(name, level, parent_lvl) (select '其他',3,id from t_entry where name = '行车交通' and level = 2);

insert into t_entry(name, level, parent_lvl) (select '电子数码',2,id from t_entry where name = '支出类' and level = 1);
insert into t_entry(name, level, parent_lvl) (select '电子产品',3,id from t_entry where name = '电子数码' and level = 2);
insert into t_entry(name, level, parent_lvl) (select '数码周边',3,id from t_entry where name = '电子数码' and level = 2);

insert into t_entry(name, level, parent_lvl) (select '虚拟产品',2,id from t_entry where name = '支出类' and level = 1);
insert into t_entry(name, level, parent_lvl) (select 'App',3,id from t_entry where name = '虚拟产品' and level = 2);
insert into t_entry(name, level, parent_lvl) (select '游戏',3,id from t_entry where name = '虚拟产品' and level = 2);
insert into t_entry(name, level, parent_lvl) (select '电影',3,id from t_entry where name = '虚拟产品' and level = 2);
insert into t_entry(name, level, parent_lvl) (select '音乐',3,id from t_entry where name = '虚拟产品' and level = 2);

insert into t_entry(name, level, parent_lvl) (select '医疗保健',2,id from t_entry where name = '支出类' and level = 1);
insert into t_entry(name, level, parent_lvl) (select '药品',3,id from t_entry where name = '医疗保健' and level = 2);
insert into t_entry(name, level, parent_lvl) (select '牙医',3,id from t_entry where name = '医疗保健' and level = 2);
insert into t_entry(name, level, parent_lvl) (select '医疗',3,id from t_entry where name = '医疗保健' and level = 2);
insert into t_entry(name, level, parent_lvl) (select '个人护理',3,id from t_entry where name = '医疗保健' and level = 2);
insert into t_entry(name, level, parent_lvl) (select '健身体育',3,id from t_entry where name = '医疗保健' and level = 2);
insert into t_entry(name, level, parent_lvl) (select '眼部护理',3,id from t_entry where name = '医疗保健' and level = 2);
insert into t_entry(name, level, parent_lvl) (select '理发',3,id from t_entry where name = '医疗保健' and level = 2);

insert into t_entry(name, level, parent_lvl) (select '衣物服饰',2,id from t_entry where name = '支出类' and level = 1);
insert into t_entry(name, level, parent_lvl) (select '配饰',3,id from t_entry where name = '衣物服饰' and level = 2);
insert into t_entry(name, level, parent_lvl) (select '服装',3,id from t_entry where name = '衣物服饰' and level = 2);
insert into t_entry(name, level, parent_lvl) (select '鞋',3,id from t_entry where name = '衣物服饰' and level = 2);

insert into t_entry(name, level, parent_lvl) (select '娱乐学习',2,id from t_entry where name = '支出类' and level = 1);
insert into t_entry(name, level, parent_lvl) (select '电影',3,id from t_entry where name = '娱乐学习' and level = 2);
insert into t_entry(name, level, parent_lvl) (select 'KTV',3,id from t_entry where name = '娱乐学习' and level = 2);
insert into t_entry(name, level, parent_lvl) (select '麻将扑克',3,id from t_entry where name = '娱乐学习' and level = 2);
insert into t_entry(name, level, parent_lvl) (select '图书',3,id from t_entry where name = '娱乐学习' and level = 2);

insert into t_entry(name, level, parent_lvl) (select '家居生活',2,id from t_entry where name = '支出类' and level = 1);
insert into t_entry(name, level, parent_lvl) (select '家电',3,id from t_entry where name = '家居生活' and level = 2);
insert into t_entry(name, level, parent_lvl) (select '家居家装',3,id from t_entry where name = '家居生活' and level = 2);
insert into t_entry(name, level, parent_lvl) (select '快递费',3,id from t_entry where name = '家居生活' and level = 2);
insert into t_entry(name, level, parent_lvl) (select '维修费',3,id from t_entry where name = '家居生活' and level = 2);
insert into t_entry(name, level, parent_lvl) (select '生活用品',3,id from t_entry where name = '家居生活' and level = 2);

insert into t_entry(name, level, parent_lvl) (select '其他杂项',2,id from t_entry where name = '支出类' and level = 1);
insert into t_entry(name, level, parent_lvl) (select '丢失',3,id from t_entry where name = '其他杂项' and level = 2);
insert into t_entry(name, level, parent_lvl) (select '坏账',3,id from t_entry where name = '其他杂项' and level = 2);
