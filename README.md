# IdGenerate
generate id for distributed system

与雪花算法不同，采用的是数据库/redis或者其他介质来存储已生成的id，并保证可读性、唯一性。

#基本思路
##可读性：
使用指定的前缀/后缀生成
##唯一性
连接数据库/redis，采用数据库的唯一索引方式、redis事务的方式，保证数据唯一