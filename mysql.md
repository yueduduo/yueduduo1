# MySQLstudyday1
### 关系型数据库
![rdbms](RDBMS2022-03-06%20164836.png)

### 启动关闭Mysql服务
* net start mysql
* net stop mysql
  
### 连接Mysql服务
mysql [-h 主机] [-P 端口] -u username -p[密码]
* -p后无空格
* 若没写-h 主机，默认本机localhost
* 若没写-P 端口，默认3306

***
### SQL通用语法
* 语句可单行可可多行，**以分号结尾**
* 不区分大小写，关键字尽量大写
* 注释，单行--或#（仅MySQL），多行/*... */
### SQL语句分类
* DDL，数据定义语言，定义数据库对象,如create表、库、字段等
* DML，数据操作语句，insert增、delete删、update改
* DQL，数据查询语句，如select查表
* DCL，数据控制语句，创数据库用户、控制权限，如grant授权，revoke撤权

***

### DDL
* 创建 CREATE DATABASE [IF NOT EXISTS] 数据库名 [DEFAULT CHARSET charser_name] [COLLATE collation_name];
  ![](creat2022-03-06%20172538.png)

* 查看
  * 看所有数据库 SHOW DATABASES;
  * 看当前数据库 SELECT DATABASE();
  * 看之前创的数据库data_name的定义信息 SHOW CREATE DATABASE data_name;

* 【慎重】删 DROP DATABASE [IF EXISTS]数据库名;

* 使用 USE 数据库名;