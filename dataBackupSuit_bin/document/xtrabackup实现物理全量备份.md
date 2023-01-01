# xtrabackup实现物理全量备份

### 一、在源库创建备份用户

```
create user 'bkpuser'@'%' identified with mysql_native_password by 'abc123';
grant BACKUP_ADMIN, PROCESS, RELOAD, CREATE TABLESPACE, SUPER, LOCK TABLES, REPLICATION CLIENT  on *.* to bkpuser@'%' ;
GRANT SELECT ON performance_schema.log_status TO 'bkpuser'@'%';
FLUSH PRIVILEGES;
--创建用户并授予权限
```



### 二、创建备份目录和恢复目录

执行中先把源库的内容备份到“备份目录”，恢复时要先从“备份目录”恢复到“恢复目录”，然后新库的data从“恢复目录”启动。

```
--备份目录
/usr/local/mysql/data_backup
--恢复目录
/usr/local/mysql/data_recovery
```

### 三、执行备份操作

在源库正常运行状态执行xtrabackup的全量物理备份：

```
./xtrabackup --backup --target-dir=/usr/local/mysql/data_backup --user=bkpuser --password=abc123 --socket=/usr/local/mysql/data13000/mysql.sock

./xtrabackup --prepare --target-dir=/usr/local/mysql/data_backup
```

### 四、新库恢复

```
./xtrabackup --copy-back --target-dir=/usr/local/mysql/data_backup --datadir=/usr/local/mysql/data_recovery
```

### 五、恢复状态启动

停止源库的运行。

首先要恢复一个启动新库所使用的conf文件，最简单的就是复制之前源库的conf文集，修改datadir为要恢复的data目录。

之后要chown -R mysql.mysql /恢复目录，把该目录下所有的文集权属改为mysql数据库用户可用。

最后执行新库的启动，即可成功。



### 六、增量备份。

[xtrabackup增量备份参考]: https://blog.csdn.net/wzq756984/article/details/86528946

