//操作mysql5数据库
package database

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

//dataservers表结构体定义
type Dataservers struct {
	 servername string
	 host string
	 port string
	 username string
	 status string
	 master_online_status string
	 master_backup string
	 remote_user string
	 remote_port string
	 max_needed_conn_max_mysql_conn string
	 master_priority string
}

//struct of datasource
type Datasource struct {
	Name string
	Type string
	SourceStatus string
	LoadBalanceStrategy string
	MasterSource string
	SlaveSourceStatus string
	MasterServer string
	MasterServerHostInfo string
	MasterServerStatus string
	SelectMasterTime string
	SlaveSource string
	SlaveServers string
	SlaveServerHostInfo string
	SlaveServerStatus string
}

type Resultset struct {
	Count string
}


func GrantAllToRoot(DB *sql.DB){
	err := DB.QueryRow("grant all on *.* to root@'%';")
	if err != nil{
		 fmt.Printf("grant all on *.* to root@'%' :%v\n", err)
       return
	}
	 fmt.Println("grant all on *.* to root@'%' successd:")
}

func NewGlobalDatasource(DB *sql.DB){
	err := DB.QueryRow("dbscale dynamic add schema dataspace global_db datasource = global_ds;")
	if err != nil{
		 fmt.Printf("dbscale dynamic add schema dataspace global_db datasource = global_ds;", err)
       return
	}
	 fmt.Println("dbscale dynamic add schema dataspace global_db datasource = global_ds; successd")
}

//获取初始正常节点数量
func FetchHealthWorkingNum(DB *sql.DB) int {
	fmt.Println("search status from dataservers")
	rows,err := DB.Query("dbscale show datasource type = replication where SlaveSourceStatus like  'Working';")
//	var count = int
	var count = 0
	if err != nil{
		fmt.Println(err)
	}
	defer 	rows.Close()
	
	for rows.Next(){
		var source Datasource
		rows.Scan(&source.Name,&source.Type,&source.SourceStatus,&source.LoadBalanceStrategy,&source.MasterSource,&source.SlaveSourceStatus,&source.MasterServer,&source.MasterServerHostInfo,&source.MasterServerStatus,&source.SelectMasterTime,&source.SlaveSource,&source.SlaveSourceStatus,&source.SlaveServers,&source.SlaveServerHostInfo,&source.SlaveServerStatus)
//	&dataserver.servername,&dataserver.host,&dataserver.port,&dataserver.username,&dataserver.status,&dataserver.master_online_status,&dataserver.master_backup,&dataserver.remote_user,&dataserver.remote_port,&dataserver.max_needed_conn_max_mysql_conn,&dataserver.master_priority
//		fmt.Println(source.Name+" "+source.MasterServerHostInfo+" "+source.SlaveSourceStatus)
		if source.SlaveSourceStatus == "Working" {
			count = count + 1
		}
	}
	return count
}

//获取表的记录数
func FetchTbCount(DB *sql.DB) string {
	fmt.Println("select count(*) from testdb.sbtest1")
	rows,err := DB.Query("select count(*) from testdb.sbtest1;")
	var result string

	if err != nil{
		fmt.Println(err)
	}
	defer 	rows.Close()
	
	for rows.Next(){
		var rs Resultset
		rows.Scan(&rs.Count)
		result = rs.Count

	}
	return result
}

//获取当前正常节点数量
func FetchCurrentWorkingNum(DB *sql.DB) int {
	fmt.Println("search status from dataservers")
	rows,err := DB.Query("dbscale show datasource type = replication where SlaveSourceStatus like  'Working';")
//	var count = int
	var count = 0
	if err != nil{
		fmt.Println(err)
	}
	defer 	rows.Close()
	
	for rows.Next(){
		var source Datasource
		rows.Scan(&source.Name,&source.Type,&source.SourceStatus,&source.LoadBalanceStrategy,&source.MasterSource,&source.SlaveSourceStatus,&source.MasterServer,&source.MasterServerHostInfo,&source.MasterServerStatus,&source.SelectMasterTime,&source.SlaveSource,&source.SlaveSourceStatus,&source.SlaveServers,&source.SlaveServerHostInfo,&source.SlaveServerStatus)
//	&dataserver.servername,&dataserver.host,&dataserver.port,&dataserver.username,&dataserver.status,&dataserver.master_online_status,&dataserver.master_backup,&dataserver.remote_user,&dataserver.remote_port,&dataserver.max_needed_conn_max_mysql_conn,&dataserver.master_priority
//		fmt.Println(source.Name+" "+source.MasterServerHostInfo+" "+source.SlaveSourceStatus)
		if source.SlaveSourceStatus == "Working" {
			count = count + 1
		}
	}
	return count
}

func FetchDataserversInfo (DB *sql.DB) bool {
	fmt.Println("search status from dataservers")
	rows,err := DB.Query("dbscale show datasource type = replication where SlaveSourceStatus like  'Working';")
	var result = false
//	var count = 0
	if err != nil{
		fmt.Println(err)
	}
	defer 	rows.Close()
	
	for rows.Next(){
		var source Datasource
		rows.Scan(&source.Name,&source.Type,&source.SourceStatus,&source.LoadBalanceStrategy,&source.MasterSource,&source.SlaveSourceStatus,&source.MasterServer,&source.MasterServerHostInfo,&source.MasterServerStatus,&source.SelectMasterTime,&source.SlaveSource,&source.SlaveSourceStatus,&source.SlaveServers,&source.SlaveServerHostInfo,&source.SlaveServerStatus)
//	&dataserver.servername,&dataserver.host,&dataserver.port,&dataserver.username,&dataserver.status,&dataserver.master_online_status,&dataserver.master_backup,&dataserver.remote_user,&dataserver.remote_port,&dataserver.max_needed_conn_max_mysql_conn,&dataserver.master_priority
//		fmt.Println(source.Name+" "+source.MasterServerHostInfo+" "+source.SlaveSourceStatus)
		if source.SlaveSourceStatus == "Working" {
			result = true
		}else{
			result = false
		}
	}
	return result
}

//获取要断网的master节点ip
func IfdownHost(DB *sql.DB,port string,role string) string {
	rows,err := DB.Query("dbscale show dataservers where port like '%"+port+"%';")
	//根据端口号信息，得到某一个元素的所有主从节点列表（在不同节点上，相同端口对应同一个角色）
	var ifdown string
	if err != nil{
		fmt.Println(err)
	}
	defer 	rows.Close()
	
	for rows.Next(){
		var dataserver Dataservers
		rows.Scan(&dataserver.servername,&dataserver.host,&dataserver.port,&dataserver.username,&dataserver.status,&dataserver.master_online_status,&dataserver.master_backup,&dataserver.remote_user,&dataserver.remote_port,&dataserver.max_needed_conn_max_mysql_conn,&dataserver.master_priority)
//		fmt.Println("servername: "+dataserver.servername+" host: "+dataserver.host+" port: "+dataserver.port+"  username: "+dataserver.username+" status: "+dataserver.status+" master_online_status: "+dataserver.master_online_status)
//		if ((dataserver.port == port)&&(dataserver.master_online_status == 'Master_Online')){
//			killed = dataserver.host
//		}
		if role == "master" {
			if ((dataserver.port == port)&&(dataserver.master_online_status == "Master_Online")) {
				ifdown = dataserver.host
			}
		}
		if role == "slave" {
			if ((dataserver.port == port)&&(dataserver.master_online_status == "Not_A_Master")) { //无论是几个slave，最后返回的肯定是最后一个slave的host.
				ifdown = dataserver.host
			}
		}
		
		//当前方法是要根据输入的端口号同时对应的master_online_status，获得master角色的host信息
	}
	return ifdown
}


func Killdataserversmaster (DB *sql.DB,port string) string {
	rows,err := DB.Query("dbscale show dataservers where status like '%normal%';")
	var killed string
	if err != nil{
		fmt.Println(err)
	}
	defer 	rows.Close()
	
	for rows.Next(){
		var dataserver Dataservers
		rows.Scan(&dataserver.servername,&dataserver.host,&dataserver.port,&dataserver.username,&dataserver.status,&dataserver.master_online_status,&dataserver.master_backup,&dataserver.remote_user,&dataserver.remote_port,&dataserver.max_needed_conn_max_mysql_conn,&dataserver.master_priority)
//		fmt.Println("servername: "+dataserver.servername+" host: "+dataserver.host+" port: "+dataserver.port+"  username: "+dataserver.username+" status: "+dataserver.status+" master_online_status: "+dataserver.master_online_status)
//		if ((dataserver.port == port)&&(dataserver.master_online_status == 'Master_Online')){
//			killed = dataserver.host
//		}
		if ((dataserver.port == port)&&(dataserver.master_online_status == "Master_Online")) {
				killed = dataserver.host
			}
		
	}
	return killed
}

func Killdataserverslastslave (DB *sql.DB,port string) string {
	rows,err := DB.Query("dbscale show dataservers where status like '%normal%';")
	var killed string
	if err != nil{
		fmt.Println(err)
	}
	defer 	rows.Close()
	
	for rows.Next(){
		var dataserver Dataservers
		
//		var slavelist []string = make([]string,1)
		rows.Scan(&dataserver.servername,&dataserver.host,&dataserver.port,&dataserver.username,&dataserver.status,&dataserver.master_online_status,&dataserver.master_backup,&dataserver.remote_user,&dataserver.remote_port,&dataserver.max_needed_conn_max_mysql_conn,&dataserver.master_priority)
		if ((dataserver.port == port)&&(dataserver.master_online_status == "Not_A_Master")) {
				killed = dataserver.host
			}
	}
	return killed
}

func DisableServerInfo (DB *sql.DB,servername string,rolename string) []string {
	rows,err := DB.Query("dbscale show dataservers where servername like '%"+servername+"%';")
	var disableInfo []string = make([]string,0)
//	var disableInfo string
	if err != nil{
		fmt.Println(err)
	}
	defer 	rows.Close()
	
	for rows.Next(){
		var dataserver Dataservers
		rows.Scan(&dataserver.servername,&dataserver.host,&dataserver.port,&dataserver.username,&dataserver.status,&dataserver.master_online_status,&dataserver.master_backup,&dataserver.remote_user,&dataserver.remote_port,&dataserver.max_needed_conn_max_mysql_conn,&dataserver.master_priority)
//		disablelist = append(disablelist,dataserver.servername+","+dataserver.host+","+dataserver.port+","+dataserver.master_online_status)
		if (rolename == "master")&&(dataserver.master_online_status == "Master_Online") {
			disableInfo = append(disableInfo,dataserver.servername+","+dataserver.host+","+dataserver.port)
		}
		if (rolename == "slave")&&(dataserver.master_online_status == "Not_A_Master") {
			disableInfo = append(disableInfo,dataserver.servername+","+dataserver.host+","+dataserver.port)
		}
		
	}
	return disableInfo
}

func ChangemasterServerInfo (DB *sql.DB,cmport string) []string {
	rows,err := DB.Query("dbscale show dataservers where port like  '%"+cmport+"%';")
	var changemasterInfo []string = make([]string,0)	
//	var disableInfo string
	if err != nil{
		fmt.Println(err)
	}
	defer 	rows.Close()
	
	for rows.Next(){
		var dataserver Dataservers
		rows.Scan(&dataserver.servername,&dataserver.host,&dataserver.port,&dataserver.username,&dataserver.status,&dataserver.master_online_status,&dataserver.master_backup,&dataserver.remote_user,&dataserver.remote_port,&dataserver.max_needed_conn_max_mysql_conn,&dataserver.master_priority)
		if (dataserver.master_online_status == "Not_A_Master") {
			changemasterInfo = append(changemasterInfo,dataserver.servername+","+dataserver.host)
		}
		
	}
	return changemasterInfo
}

