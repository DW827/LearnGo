# 关系型数据库
+ 数据库结构：client->cmd->mysql管理->database,DB数据库->table,数据表->数据单例
### 一，如何使用终端操作数据库
+ 登录数据库：mysql -uroot -p123456
+ 退出数据库服务器：exit
+ 查询数据库服务器中所有的数据库：show databases;
+ 进入某一个数据库，如进入giligili库：use giligili;
+ 查看该数据库的所有表：show tables;
```
+--------------------+
| Tables_in_giligili |
+--------------------+
| users              |
| videos             |
+--------------------+
2 rows in set (0.00 sec)
```
+ 一系列sql语句，INSERT增，DELETE删，UPDATE改，SELECT查，如查询语句：select * from videos;
```
+----+-------------------------+-------------------------+------------+-----------------------------+-----------------------------+------+--------+---------+
| id | created_at              | updated_at              | deleted_at | title                       | info                        | url  | avatar | user_id |
+----+-------------------------+-------------------------+------------+-----------------------------+-----------------------------+------+--------+---------+
|  1 | 2022-01-14 11:55:42.805 | 2022-01-14 13:33:14.219 | NULL       | 咸鱼投手如何逆袭            | 咸鱼投手如何逆袭            |      |        |       1 |
|  2 | 2022-01-14 13:33:33.003 | 2022-01-14 13:34:17.012 | NULL       | 王牌投手的养成之路          | 王牌投手的养成之路          |      |        |       1 |
+----+-------------------------+-------------------------+------------+-----------------------------+-----------------------------+------+--------+---------+
2 rows in set (0.00 sec)
```
#### 数据库操作创建删除，数据记录操作语句增删查改
+ 下面演示一下简单常用的操作：“新建数据库-->新建表（数据定义）-->插入信息（增）-->查询信息（查）-->更新信息（改）-->删除信息（删）-->删除表-->删除数据库”
```
// 1.新建数据库:
mysql> create database new_database;
Query OK, 1 row affected (0.50 sec)
// 查看库
mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| giligili           |
| information_schema |
| mysql              |
| new_database       |
| performance_schema |
| sys                |
+--------------------+
6 rows in set (0.00 sec)

// 2.新建表（数据定义）：
mysql> use new_database;
Database changed
mysql> create table new_table(
    -> id int,
    -> usename char(10));
Query OK, 0 rows affected (0.47 sec)
// 查看表
mysql> show tables;
+------------------------+
| Tables_in_new_database |
+------------------------+
| new_table              |
+------------------------+
1 row in set (0.00 sec)

// 3.插入数据（增）：
mysql>  insert into new_table(id,username) values(1,'tom');
ERROR 1054 (42S22): Unknown column 'username' in 'field list'
mysql>  insert into new_table(id,usename) values(1,'tom');
Query OK, 1 row affected (0.38 sec)

// 4.查询数据（查）：
mysql> select * from new_table;
+------+---------+
| id   | usename |
+------+---------+
|    1 | tom     |
+------+---------+
1 row in set (0.00 sec)

// 5.更新信息（改）：
mysql> update new_table
    -> set id=2,usename="cat"
    -> where id=1
    -> ;
Query OK, 1 row affected (0.35 sec)
Rows matched: 1  Changed: 1  Warnings: 0

// 6.删除信息（删）：
mysql> delete from new_table
    -> where id=2
    -> ;
Query OK, 1 row affected (0.10 sec)

mysql> select * from new_table;
Empty set (0.00 sec)

// 7.删除表：
mysql> drop table new_table;
Query OK, 0 rows affected (0.13 sec)

mysql> show tables;
Empty set (0.00 sec)

// 8.删除数据库：
mysql> drop database new_database;
Query OK, 0 rows affected (0.19 sec)

mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| giligili           |
| information_schema |
| mysql              |
| performance_schema |
| sys                |
+--------------------+
5 rows in set (0.00 sec)
```
#### WHERE子句，LIKE子句，[UNION语句](https://www.runoob.com/mysql/mysql-union-operation.html)
+ 在 MySQL 中使用 SELECT 命令来读取数据，同时可以在 SELECT 语句中使用 WHERE 子句来获取指定的记录。使用等号 = 来设定获取数据的条件，如 ：
```
SELECT * from table_name WHERE field1 = '%somevalue'
```
+ LIKE 子句中使用百分号 %字符来表示任意字符,字段含有 "somevalue" 字符的所有记录,可以在SELECT, DELETE 或 UPDATE 命令中使用 WHERE...LIKE 子句来指定条件。以下是 SQL SELECT 语句使用 LIKE 子句从数据表中读取数据的通用语法：
```
SELECT * from table_name WHERE field1 LIKE '%somevalue'
```
+ UNION 语句：用于将不同表中相同列中查询的数据结果组合到一个结果集合中展示出来；（不包括重复数据）
+ UNION ALL 语句：用于将不同表中相同列中查询的数据结果组合到一个结果集合中展示出来；（包括重复数据）
使用形式如下：
```
SELECT 列名称 FROM 表名称 UNION SELECT 列名称 FROM 表名称 ORDER BY 列名称；
SELECT 列名称 FROM 表名称 UNION ALL SELECT 列名称 FROM 表名称 ORDER BY 列名称；
```
#### MySQL排序
+ 使用 MySQL 的 ORDER BY 子句来设定你想按哪个字段哪种方式来进行排序，再返回搜索结果。
```
SELECT field1, field2,...fieldN FROM table_name1, table_name2...
ORDER BY field1 [ASC [DESC][默认 ASC]], [field2...] [ASC [DESC][默认 ASC]]
```
#### MySQL[分组](https://www.runoob.com/mysql/mysql-group-by-statement.html)
GROUP BY 语句根据一个或多个列对结果集进行分组。
在分组的列上我们可以使用 COUNT, SUM, AVG,等函数。
```
SELECT column_name, function(column_name)
FROM table_name
WHERE column_name operator value
GROUP BY column_name;
```

#### MySQL[连接](https://www.runoob.com/mysql/mysql-join.html)
在 SELECT, UPDATE 和 DELETE 语句中使用 Mysql 的 JOIN 来联合多表查询。
```
mysql> SELECT a.runoob_id, a.runoob_author, b.runoob_count FROM runoob_tbl a INNER JOIN tcount_tbl b ON a.runoob_author = b.runoob_author;
+-------------+-----------------+----------------+
| a.runoob_id | a.runoob_author | b.runoob_count |
+-------------+-----------------+----------------+
| 1           | 菜鸟教程    | 10             |
| 2           | 菜鸟教程    | 10             |
| 3           | RUNOOB.COM      | 20             |
| 4           | RUNOOB.COM      | 20             |
+-------------+-----------------+----------------+
4 rows in set (0.00 sec)
```

#### MySQL[NULL处理](https://www.runoob.com/mysql/mysql-null.html)
MySQL提供了三大运算符:
IS NULL: 当列的值是 NULL,此运算符返回 true。
IS NOT NULL: 当列的值不为 NULL, 运算符返回 true。
<=>: 比较操作符（不同于 = 运算符），当比较的的两个值相等或者都为 NULL 时返回 true。
关于 NULL 的条件比较运算是比较特殊的。你不能使用 = NULL 或 != NULL 在列中查找 NULL 值 。

#### MySQL [数据类型](https://www.runoob.com/mysql/mysql-data-types.html)
+ MySQL 支持多种类型，大致可以分为三类：数值，日期/时间，字符串(字符)类型

#### MySQL建表约束
+ --主键约束，field primary key 或 primary(field)：能够唯一确定一张表中的一条记录，也就是通过给某个字段添加约束使得该字段不重复且不为空。
如：不同用户不同ID，但是同名。
```
mysql> create table user(
    -> id int primary key,
    -> name varchar(20));
Query OK, 0 rows affected (0.47 sec)

mysql> desc user;
+-------+-------------+------+-----+---------+-------+
| Field | Type        | Null | Key | Default | Extra |
+-------+-------------+------+-----+---------+-------+
| id    | int(11)     | NO   | PRI | NULL    |       |
| name  | varchar(20) | YES  |     | NULL    |       |
+-------+-------------+------+-----+---------+-------+
2 rows in set (0.00 sec)

mysql> insert into user values(null,'zhangsan');
ERROR 1048 (23000): Column 'id' cannot be null

mysql> insert into user values(1,'zhangsan');
Query OK, 1 row affected (0.09 sec)

mysql> insert into user values(1,'zhangsan');
ERROR 1062 (23000): Duplicate entry '1' for key 'PRIMARY'

mysql> insert into user values(2,'zhangsan');
Query OK, 1 row affected (0.10 sec)

mysql> select * from user;
+----+----------+
| id | name     |
+----+----------+
|  1 | zhangsan |
|  2 | zhangsan |
+----+----------+
2 rows in set (0.00 sec)
```
+ -- 联合主键约束，primary key(field1, field2)：联合的主键不同时重复，任一主键不为空。
如：多对多联系如何唯一标识一个元组，因为单靠学号或单靠课程号都无法唯一标示一个元组（因为每个学生可以选修多门课程，每门课程可以有多名学生选修）
```
mysql> create table user1(
    -> id int,
    -> name varchar(20),
    -> primary(id,name)
    -> );
```
+ --修改表的主键约束
> 添加主键：alter table table_name add primary key(field);
> 删除主键：alter table table_name drop primary key;
> 修改字段同时添加主键：alter table table_name modify field field_type primary key;
###### [alter](https://www.runoob.com/mysql/mysql-alter.html) 修改数据表名或者删除DROP，添加ADD，修改数据表字段
> alter修改表名ALTER TABLE old_table RENAME TO new_table;
> modify修改字段类型
> change修改字短名称+类型
+ --自增约束，primary key auto_increment: 和主键约束搭配使用，创建一条记录时会自动生成不重复字段。如：
```
mysql> create table user2(
    -> id int primary key auto_increment,
    -> name varchar(20));
Query OK, 0 rows affected (0.53 sec)

mysql> insert into user2(name) values ('zhangsan');
Query OK, 1 row affected (0.39 sec)
mysql> insert into user2(name) values ('zhangsan');
Query OK, 1 row affected (0.39 sec)

mysql> select * from user2;
+----+----------+
| id | name     |
+----+----------+
|  1 | zhangsan |
|  2 | zhangsan |
+----+----------+
2 rows in set (0.00 sec)
```
+ --唯一约束，unique：修饰某一字段不能重复。唯一约束字段可以为空，区别于主键约束，一个表只有一个主键约束，主键非空且唯一。复合主键：unique(field1,field2,...)，类似于联合主键，主键不同时重复，但是可以为空。
```
mysql> create table user3(
    -> id int,
    -> name varchar(20),
    -> unique(id,name)      // 类似于联合主键
    -> );
Query OK, 0 rows affected (0.22 sec)

mysql> insert into user3 values(1,'zhangsan');
Query OK, 1 row affected (0.11 sec)

mysql> insert into user3 values(2,'zhangsan');
Query OK, 1 row affected (0.10 sec)

mysql> select * from user3;
+------+----------+
| id   | name     |
+------+----------+
|    1 | zhangsan |
|    2 | zhangsan |
+------+----------+
2 rows in set (0.02 sec)
```
+ --非空约束，not null：修饰的字段不能为空。
+ --默认约束，default：修饰的字段在插入记录时如果没有给字段赋值会自动使用默认值。
+ --外键约束，foreign key(...) references ...
> 副表（子表）只能使用主表中有的数据值，不能修改主表（父表)。
```
mysql> create table classes(
    -> id int primary key,
    -> name varchar(20)
    -> );
Query OK, 0 rows affected (0.58 sec)

mysql> create table students(
    -> id int primary key,
    -> name varchar(20),
    -> class_id int,
    -> foreign key(class_id) references classes(id)
    -> );
Query OK, 0 rows affected (0.53 sec)
```
#### 数据库的三大范式NF
+ 第一范式：拆字段直到不能拆为止。（如：国家省市区镇街道号）
+ 第二范式：满足第一范式前提下，除主键外的每一列必须完全依赖于主键。
如果要出现不完全依赖，只可能发生在联合主键的情况下。
如果建表除主键外的其他列，只依赖于联合主键的部分字段，那么要拆表。
好处：可以通过主键找到其他的数据，主表从表。
+ 第三范式：满足第二范式的前提下，除主键外的其他列之间不能有传递依赖关系。

#### MySQL查询练习
+ 1. 查询table1的所有数据记录
select * from table1;
+ 2. 查询table1的field1和field2列的记录
select field1， field2 from table1；
+ 3. 查询table1的不重复的field1列，即去重distinct
select distinct field1 from table1;
+ 4. 查询table1的field1列中值为60~80之间的所有记录
select * from table1 where field1 between 60 and 80;
select * from table1 where field1 > 60 and field < 80;
+ 5. 查询table1的field列值为60，80，或88的所有记录
select * from table1 where field1 in(60,80,88);
+ 6. 查询table1的field1=“v1”或field2=“v2”的所有记录
select * from table1 where field1="v1" or field2="v2";
+ 7. 查询table1的以field1降序结果返回的所有记录
select * from table1 order by field1 desc;
+ 8. 查询table1的以field1降序，field2升序
select * from table1 order by field1 asc, order by field2 desc;
+ 9. 统计table1的field1=“v1”的记录条数
select count(*) from table1 where field1="v1";
+ 10. 查询table1中field1值最高的field2和field3列
select field2,field3 from table1 where field1=(select max(field1) from table1);
select field2,field3 from table1 order by field1 desc limit 0,1;
注:limit offset,start
+ 11. 查询table1的每个不重复field1列的平均field2值
select field1,avg(field2) from table1 group by field1;
+ 12. 查询table1的至少有两个记录并以3开头的field1的平均field2值,(分组group by 和 having 搭配)
select field1,avg(field2) from table1 group by field1 having count(field1)>=2 and like'3%';
+ 13. 查询table1的field1值大于70小于90的field2列
select field2, field1 from table1 where filed1>70 and field1<90;
select field2, field1 from table1 where field1 between 70 and 90;
+ 14. 多表查询，查询所有table1的field1列，table2的field2，field3列，（其中table2的field4外键约束于table1的field4）
select field1, field2, field3 from table1, table2 where table1.field4=table2.field4;
+ 15. 复杂的子表查询时，分段查询再整合到一个命令。
+ 16. ＞ any(values) 大于至少一个
+ 17. ＞ all(values) 大于任意一个

