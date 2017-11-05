# database

## 账户

资产类和负债类会计科目有账户。收入类和支出类没有实体账户，虚拟。

表名：t_account

<table>
<tr>
    <th>字段</th>
    <th>类型</th>
    <th>描述</th>
</tr>

<tr>
<td>id</td>
<td>int</td>
<td>自增主键</td>
</tr>
<tr>
<td>name</td>
<td>varchar(255)</td>
<td>账户名称</td>
</tr>
<tr>
<td>entry_id</td>
<td>int</td>
<td>会计科目，顶级科目</td>
</tr>
<tr>
<td>balance</td>
<td>int</td>
<td>当前余额</td>
</tr>

</table>

## 借贷流水

表名：t_record

<table>
<tr>
    <th>字段</th>
    <th>类型</th>
    <th>描述</th>
</tr>

<tr>
<td>id</td>
<td>int</td>
<td>自增主键</td>
</tr>
<tr>
<td>type</td>
<td>int</td>
<td>借贷类型，1-借，2-贷</td>
</tr>
<tr>
<td>account_id</td>
<td>int</td>
<td>借贷账户</td>
</tr>
<tr>
<td>entry_id</td>
<td>int</td>
<td>会计科目</td>
</tr>
<tr>
<td>amount</td>
<td>int</td>
<td>金额</td>
</tr>
<tr>
<td>datetime</td>
<td>datetime</td>
<td>日期</td>
<tr>
<td>counter</td>
<td>varchar(255)</td>
<td>借贷对象，只有支出类和收入类对象有此信息</td>
</tr>
</tr>

</table>

## 会计科目

表名：t_entry

<table>
<tr>
    <th>字段</th>
    <th>类型</th>
    <th>描述</th>
</tr>

<tr>
<td>id</td>
<td>int</td>
<td>自增主键</td>
</tr>
<tr>
<td>name</td>
<td>varchar(255)</td>
<td>名称</td>
</tr>
<tr>
<td>level</td>
<td>int</td>
<td>级别</td>
</tr>
<tr>
<td>parent_lvl</td>
<td>int</td>
<td>父级别</td>
</tr>
</table>

##### 资产类
* 现金
* 银行存款
* 投资
* 应收利息
* 应收账款

##### 负债类
* 应付账款
* 应交税金

##### 收入类
* 工资收入
* 奖金
* 加班费
* 投资收益
* 其他收入

##### 支出类

* 食品酒水
    + 早餐
    + 午餐
    + 晚餐
    + 水果
    + 饮料酒水
    + 食品零食
    + 蔬菜原料
    + 油盐酱醋
    
* 账单
    + 房租
    + 水费
    + 电费
    + 燃气费
    + 管理费
    + 网络宽带
    + 话费
    + 电视费
    + VPS
    + 域名
    + 会员

* 人情关系
    + 结婚礼金
    + 红包
    + 聚餐
    + 礼物
    + 孝敬
    + 其他
    
* 行车交通
    + 公交地铁
    + 出租车
    + 客车
    + 高铁火车
    + 飞机
    
* 电子数码
    + 电子产品
    + 数码周边

* 虚拟产品
    + App
    + 游戏
    + 电影
    + 音乐

* 医疗保健
    + 药品
    + 牙医
    + 医疗
    + 个人护理
    + 健身体育
    + 眼部护理
    + 理发
    
* 衣物服饰
    + 配饰
    + 服装
    + 鞋

* 娱乐学习
    * 电影
    * KTV
    * 麻将扑克
    * 图书
    
* 家居生活
    * 卫生清洁
    * 家电
    * 家居家装
    * 快递费
    * 维修费
    * 生活用品

* 其他杂项
    * 丢失
    * 坏账
