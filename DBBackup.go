package main

import (
	"fmt"
	"strings"
	"time"
	"strconv"
	load "dataBackupSuit/load"
	remote "dataBackupSuit/remote"
	local "dataBackupSuit/local"

)

//参数：配置文件


func main(){
	fmt.Println("Database backup suit for MySQL.")
	
	var Database string
	var ClusterDb string
	var FullBackup string
	var BackupOrRestore string
	var DBUrl string
	var DBPort string
	var BkUsr string
	var BkPass string
	var BkPath string	
	var BackupPth string
	var RestorePth string
	var SshUsr string
	var SshPass string
	var DBPth string
	var DBSock string
	var DBConf string


	list := load.Loadproperties("./properties")
	for i,v := range list {
		if i == 0{ //Database:mysql5
			Database = strings.Split(v,":")[1]
//			fmt.Println("Database : "+Database)
		}
		if i == 1{ //Database:mysql5
			ClusterDb = strings.Split(v,":")[1]
//			fmt.Println("ClusterDb : "+ClusterDb)
		}
		if i == 2{ //Database:mysql5
			FullBackup = strings.Split(v,":")[1]
			
		}
		if i == 3{ //Database:mysql5
			BackupOrRestore = strings.Split(v,":")[1]
//			fmt.Println("BackupOrRestore : "+BackupOrRestore)
		}
		if i == 4{ //Database:mysql5
			DBUrl = strings.Split(v,":")[1]
		}
		if i == 5{ //Database:mysql5
			DBPort = strings.Split(v,":")[1]
//			fmt.Println("DBPort : "+DBPort)
		}
		if i == 6{ //Database:mysql5
			BkUsr = strings.Split(v,":")[1]
//			fmt.Println("BkUsr : "+BkUsr)
		}
		if i == 7{ //Database:mysql5
			BkPass = strings.Split(v,":")[1]
//			fmt.Println("BkPass : "+BkPass)
		}
		if i == 8{ //Database:mysql5
			BkPath = strings.Split(v,":")[1]
//			fmt.Println("BkPath : "+BkPath)
		}
		if i == 9{ //Database:mysql5
			BackupPth = strings.Split(v,":")[1]
//			fmt.Println("BackupPth : "+BackupPth)
		}
		if i == 10{ //Database:mysql5
			RestorePth = strings.Split(v,":")[1]
//			fmt.Println("RestorePth : "+RestorePth)
		}
		if i == 11{ //Database:mysql5
			SshUsr = strings.Split(v,":")[1]
//			fmt.Println("SshUsr : "+SshUsr)
		}
		if i == 12{ //Database:mysql5
			SshPass = strings.Split(v,":")[1]
//			fmt.Println("SshPass : "+SshPass)
		}
		if i == 13{ 
			DBPth = strings.Split(v,":")[1]
//			fmt.Println("DBPth : "+DBPth)
		}
		if i == 14{ 
			DBSock = strings.Split(v,":")[1]
//			fmt.Println("DBSock : "+DBSock)
		}
		if i == 15{ 
			DBConf = strings.Split(v,":")[1]
//			fmt.Println("DBConf : "+DBConf)
		}

	} 

			if Database == "mysql5"{
			fmt.Println("-- backup or restore mysql5 ：")
				if BackupOrRestore == "b"{//运行备份
					if FullBackup == "init"{
						fmt.Println("	--init on mysql5")
					}
					if FullBackup == "y"{
						fmt.Println("	--full backup mysql5")
						if ClusterDb == "y"{
//							--创建本次测试的备份目录 
							UrlList := strings.Split(DBUrl,",")
							timer := time.Now().Unix()
							for _,value := range UrlList{
								fmt.Println("user : "+SshUsr+"  , create backup folder on  : "+DBUrl)
								remote.Nodemission(SshUsr, SshPass, value, 22, "mkdir -p "+BkPath)
								fmt.Println("mkdir -p "+BkPath+""+BackupPth)
								remote.Nodemission(SshUsr, SshPass, value, 22, "mkdir -p "+BkPath+""+BackupPth)
								fmt.Println("mkdir -p "+BkPath+"/"+BackupPth+"/full"+value+"_"+strconv.FormatInt(timer, 10))
								remote.Nodemission(SshUsr, SshPass, value, 22, "mkdir -p "+BkPath+""+BackupPth+"/full"+value+"_"+strconv.FormatInt(timer, 10))
							}
						}
					if ClusterDb == "n"{
							fmt.Println("backup url : "+DBUrl)
							timer := time.Now().Unix()
							fmt.Println("user : "+SshUsr+"  , create backup folder on  : "+DBUrl)
							remote.Nodemission(SshUsr, SshPass, DBUrl, 22, "mkdir -p "+BkPath)
							fmt.Println("mkdir -p "+BkPath+""+BackupPth)
							remote.Nodemission(SshUsr, SshPass, DBUrl, 22, "mkdir -p "+BkPath+""+BackupPth)
						
							fmt.Println("mkdir -p "+BkPath+""+BackupPth+"/full_"+strconv.FormatInt(timer, 10))
							remote.Nodemission(SshUsr, SshPass, DBUrl, 22, "mkdir -p "+BkPath+""+BackupPth+"/full_"+strconv.FormatInt(timer, 10))
						
							fmt.Println("backup single db for full backup ...")
							fmt.Println("		set database readonly on...")
							local.Localmysql(DBPth,BkUsr, BkPass, DBUrl, DBPort, "set global read_only = on;")
							local.Localcmd("./tool/xtrabackup_mysql5.7/bin", "xtrabackup", "--backup --target-dir='"+BkPath+""+BackupPth+"/full_"+strconv.FormatInt(timer, 10)+"' --user='"+BkUsr+"' --password='"+BkPass+"' --socket='"+DBPth+DBSock+"'")
							fmt.Println("		set database readonly off...")
							local.Localmysql(DBPth,BkUsr, BkPass, DBUrl, DBPort, "set global read_only = off;")
							local.Localcmd("./tool/xtrabackup_mysql5.7/bin", "xtrabackup", "--prepare --target-dir='"+BkPath+""+BackupPth+"/full_"+strconv.FormatInt(timer, 10)+"' ")
							local.LocalCpConf(DBPth+DBConf, BkPath+""+BackupPth+"/full_"+strconv.FormatInt(timer, 10))
						}
				}
				if FullBackup == "n"{
					fmt.Println("	--increment backup mysql5")
				}
			}
			if BackupOrRestore == "r"{//运行恢复
				fmt.Println("restore db")
				if FullBackup == "y"{
					fmt.Println("	--full restore mysql5")
					if ClusterDb == "y"{
						fmt.Println("restore url : "+DBUrl+" for use")
	//						--创建本次测试的备份目录 
							UrlList := strings.Split(DBUrl,",")
							timer := time.Now().Unix()
							for _,value := range UrlList{
	
								remote.Nodemission(SshUsr, SshPass, value, 22, "mkdir -p "+BkPath)
	
								remote.Nodemission(SshUsr, SshPass, value, 22, "mkdir -p "+BkPath+""+RestorePth)
	
								remote.Nodemission(SshUsr, SshPass, value, 22, "mkdir -p "+BkPath+""+RestorePth+"/full"+value+"_"+strconv.FormatInt(timer, 10))
							}
					}
					if ClusterDb == "n"{
						fmt.Println("backup url : "+DBUrl)
						timer := time.Now().Unix()
	
							remote.Nodemission(SshUsr, SshPass, DBUrl, 22, "mkdir -p "+BkPath)
	
							remote.Nodemission(SshUsr, SshPass, DBUrl, 22, "mkdir -p "+BkPath+""+RestorePth)
					
							fmt.Println("restore single db for full restore ...")
							local.Localcmd("./tool/xtrabackup_mysql5.7/bin", "xtrabackup", "--copy-back --target-dir='"+BkPath+""+BackupPth+"' --datadir='"+BkPath+""+RestorePth+"/full_"+strconv.FormatInt(timer, 10)+"' ")
							
					}
				}
				if FullBackup == "n"{
					fmt.Println("	--increment backup mysql5")
				}
				
			}
		
		
		}
		if Database == "mysql8"{
			fmt.Println("-- 备份或还原 mysql5 数据库：")
		}
		if Database == "postgresql"{
			fmt.Println("-- 备份或还原 postgresql 数据库：")
		}
			

	
}

