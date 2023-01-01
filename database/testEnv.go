//用来在数据库上准备测试表
package database

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

//高可用环境中用到的sysbench测试库数据，testdb相关。
func CreateTestDb(DB *sql.DB) {
	fmt.Println("create database testdb;")
	sql := "create database testdb;"
	
	if _,err := DB.Exec(sql);err != nil {
		fmt.Println("create database testdb failed  ",err)
		return
	}
		fmt.Println("create table successd")
}

//创建表规则和分片规则
func CreateRule(DB *sql.DB) {
	fmt.Println("create rule on datasource;")
	sql1 := "dbscale dynamic add normal_table dataspace testdb.sbtest1 datasource = \"global_ds\";"
	if _,err := DB.Exec(sql1);err != nil {
		fmt.Println("dbscale dynamic add normal_table dataspace testdb.sbtest1 datasource = \"global_ds\"  failed  ",err)
		return
	}else{
		fmt.Println("dbscale dynamic add normal_table dataspace testdb.sbtest1 datasource = \"global_ds\"; successd")
	}
	
	sql2 := "dbscale dynamic add normal_table dataspace testdb.sbtest2 datasource = \"normal_0\";"
	if _,err := DB.Exec(sql2);err != nil {
		fmt.Println("dbscale dynamic add normal_table dataspace testdb.sbtest2 datasource = \"normal_0\";  failed  ",err)
		return
	}else{
		fmt.Println("dbscale dynamic add normal_table dataspace testdb.sbtest2 datasource = \"normal_0\"; successd")
	}
	
	sql3 := "dbscale dynamic add hash_type partition_scheme \"hash-mod-sh\"  partition=\"par_2\" partition=\"par_1\" partition=\"par_0\"  is_shard shard_nums 12;"
	if _,err := DB.Exec(sql3);err != nil {
		fmt.Println("dbscale dynamic add hash_type partition_scheme \"hash-mod-sh\"  partition=\"par_2\" partition=\"par_1\" partition=\"par_0\"  is_shard shard_nums 12;  failed  ",err)
		return
	}else{
		fmt.Println("dbscale dynamic add hash_type partition_scheme \"hash-mod-sh\"  partition=\"par_2\" partition=\"par_1\" partition=\"par_0\"  is_shard shard_nums 12; successd")
	}

	sql4 := "dbscale dynamic add partition_table dataspace testdb.sbtest3 partition_key=\"id\" partition_scheme=\"hash-mod-sh\";"
	if _,err := DB.Exec(sql4);err != nil {
		fmt.Println("dbscale dynamic add partition_table dataspace testdb.sbtest3 partition_key=\"id\" partition_scheme=\"hash-mod-sh\";  failed  ",err)
		return
	}else{
		fmt.Println("dbscale dynamic add partition_table dataspace testdb.sbtest3 partition_key=\"id\" partition_scheme=\"hash-mod-sh\"; successd")
	}		
}

func DropRule(DB *sql.DB){
	fmt.Println("drop rule")
//	sql1 := "drop table testdb.sbtest1;"
//	if _,err := DB.Exec(sql1);err != nil {
//		fmt.Println(" --  failed  ",err)
//		return
//	}else{
//		fmt.Println("drop table testdb.sbtest1; successd")
//	}
//	
//	sql2 := "drop table testdb.sbtest2;"
//	if _,err := DB.Exec(sql2);err != nil {
//		fmt.Println(" --  failed  ",err)
//		return
//	}else{
//		fmt.Println("drop table testdb.sbtest2; successd")
//	}
//	
//	sql3 := "drop table testdb.sbtest3;"
//	if _,err := DB.Exec(sql3);err != nil {
//		fmt.Println(" --  failed  ",err)
//		return
//	}else{
//		fmt.Println("drop table testdb.sbtest3; successd")
//	}
	
	sql4 := "dbscale dynamic remove table testdb.sbtest1;"
	if _,err := DB.Exec(sql4);err != nil {
		fmt.Println(" --  failed  ",err)
		return
	}else{
		fmt.Println("dbscale dynamic remove table testdb.sbtest1; successd")
	}
	
	sql5 := "dbscale dynamic remove table testdb.sbtest2;"
	if _,err := DB.Exec(sql5);err != nil {
		fmt.Println(" --  failed  ",err)
		return
	}else{
		fmt.Println("dbscale dynamic remove table testdb.sbtest2; successd")
	}
	
	sql6 := "dbscale dynamic remove table testdb.sbtest3;"
	if _,err := DB.Exec(sql6);err != nil {
		fmt.Println(" --  failed  ",err)
		return
	}else{
		fmt.Println("dbscale dynamic remove table testdb.sbtest3; successd")
	}
	
	sql7 := "dbscale dynamic remove partition_scheme hash-mod-sh;"
	if _,err := DB.Exec(sql7);err != nil {
		fmt.Println(" --  failed  ",err)
		return
	}else{
		fmt.Println("dbscale dynamic remove partition_scheme hash-mod-sh; successd")
	}
	
	sql8 := "dbscale dynamic remove schema testdb;"
	if _,err := DB.Exec(sql8);err != nil {
		fmt.Println(" --  failed  ",err)
		return
	}else{
		fmt.Println("dbscale dynamic remove schema testdb; successd")
	}
			
}

func DropTestDb(DB *sql.DB) {
	fmt.Println("drop database testdb;")
	sql := "drop database testdb;"
	
	if _,err := DB.Exec(sql);err != nil {
		fmt.Println("drop database testdb failed  ",err)
		return
	}
		fmt.Println("drop database successd")
}
//高可用环境中用到的sysbench测试库数据，testdb相关。


//性能测试用到的sysbench测试库数据，stressdb相关。
func CreateStressDb(DB *sql.DB) {
	fmt.Println("create database stressdb;")
	sql := "create database stressdb;"
	
	if _,err := DB.Exec(sql);err != nil {
		fmt.Println("create database stressdb failed  ",err)
		return
	}
		fmt.Println("create stressdb successd")
}

//创建表规则和分片规则
func CreateStressRule(DB *sql.DB) {
	fmt.Println("create rule on datasource;")
	sql1 := "dbscale dynamic add normal_table dataspace stressdb.sbtest1 datasource = \"global_ds\";"
	if _,err := DB.Exec(sql1);err != nil {
		fmt.Println("dbscale dynamic add normal_table dataspace stressdb.sbtest1 datasource = \"global_ds\"  failed  ",err)
		return
	}else{
		fmt.Println("dbscale dynamic add normal_table dataspace stressdb.sbtest1 datasource = \"global_ds\"; successd")
	}
	
	sql2 := "dbscale dynamic add normal_table dataspace stressdb.sbtest2 datasource = \"normal_0\";"
	if _,err := DB.Exec(sql2);err != nil {
		fmt.Println("dbscale dynamic add normal_table dataspace stressdb.sbtest2 datasource = \"normal_0\";  failed  ",err)
		return
	}else{
		fmt.Println("dbscale dynamic add normal_table dataspace stressdb.sbtest2 datasource = \"normal_0\"; successd")
	}
	
	sql3 := "dbscale dynamic add hash_type partition_scheme \"hash-mod-stress\"  partition=\"par_2\" partition=\"par_1\" partition=\"par_0\"  is_shard shard_nums 12;"
	if _,err := DB.Exec(sql3);err != nil {
		fmt.Println("dbscale dynamic add hash_type partition_scheme \"hash-mod-stress\"  partition=\"par_2\" partition=\"par_1\" partition=\"par_0\"  is_shard shard_nums 12;  failed  ",err)
		return
	}else{
		fmt.Println("dbscale dynamic add hash_type partition_scheme \"hash-mod-stress\"  partition=\"par_2\" partition=\"par_1\" partition=\"par_0\"  is_shard shard_nums 12; successd")
	}

	sql4 := "dbscale dynamic add partition_table dataspace stressdb.sbtest3 partition_key=\"id\" partition_scheme=\"hash-mod-stress\";"
	if _,err := DB.Exec(sql4);err != nil {
		fmt.Println("dbscale dynamic add partition_table dataspace stressdb.sbtest3 partition_key=\"id\" partition_scheme=\"hash-mod-stress\";  failed  ",err)
		return
	}else{
		fmt.Println("dbscale dynamic add partition_table dataspace stressdb.sbtest3 partition_key=\"id\" partition_scheme=\"hash-mod-stress\"; successd")
	}		
}

func DropStressRule(DB *sql.DB){
	fmt.Println("drop rule")
//	sql1 := "drop table stressdb.sbtest1;"
//	if _,err := DB.Exec(sql1);err != nil {
//		fmt.Println(" --  failed  ",err)
//		return
//	}else{
//		fmt.Println("drop table stressdb.sbtest1; successd")
//	}
//	
//	sql2 := "drop table stressdb.sbtest2;"
//	if _,err := DB.Exec(sql2);err != nil {
//		fmt.Println(" --  failed  ",err)
//		return
//	}else{
//		fmt.Println("drop table stressdb.sbtest2; successd")
//	}
//	
//	sql3 := "drop table stressdb.sbtest3;"
//	if _,err := DB.Exec(sql3);err != nil {
//		fmt.Println(" --  failed  ",err)
//		return
//	}else{
//		fmt.Println("drop table stressdb.sbtest3; successd")
//	}
	
	sql4 := "dbscale dynamic remove table stressdb.sbtest1;"
	if _,err := DB.Exec(sql4);err != nil {
		fmt.Println(" --  failed  ",err)
		return
	}else{
		fmt.Println("dbscale dynamic remove table stressdb.sbtest1; successd")
	}
	
	sql5 := "dbscale dynamic remove table stressdb.sbtest2;"
	if _,err := DB.Exec(sql5);err != nil {
		fmt.Println(" --  failed  ",err)
		return
	}else{
		fmt.Println("dbscale dynamic remove table stressdb.sbtest2; successd")
	}
	
	sql6 := "dbscale dynamic remove table stressdb.sbtest3;"
	if _,err := DB.Exec(sql6);err != nil {
		fmt.Println(" --  failed  ",err)
		return
	}else{
		fmt.Println("dbscale dynamic remove table stressdb.sbtest3; successd")
	}
	
	sql7 := "dbscale dynamic remove partition_scheme hash-mod-stress;"
	if _,err := DB.Exec(sql7);err != nil {
		fmt.Println(" --  failed  ",err)
		return
	}else{
		fmt.Println("dbscale dynamic remove partition_scheme hash-mod-stress; successd")
	}
	
	sql8 := "dbscale dynamic remove schema stressdb;"
	if _,err := DB.Exec(sql8);err != nil {
		fmt.Println(" --  failed  ",err)
		return
	}else{
		fmt.Println("dbscale dynamic remove schema stressdb; successd")
	}
			
}

func DropStressDb(DB *sql.DB) {
	fmt.Println("drop database stressdb;")
	sql := "drop database stressdb;"
	
	if _,err := DB.Exec(sql);err != nil {
		fmt.Println("drop database stressdb failed  ",err)
		return
	}
		fmt.Println("drop database successd")
}
//性能测试用到的sysbench测试库数据，stressdb相关。

