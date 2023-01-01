package load

import (
	 "os"
    "log"
    "bufio"
    "strings"	
    "fmt"
)

func Loadproperties(profile string) []string{
	fmt.Println("读取配置文件的go")
	var sliceinfo []string
	sliceinfo,err := HandleTextProperties(profile)
    if err != nil {
        panic(err)
    }else{
    	//没有错误时，直接把切片返回个main函数
		return sliceinfo
    }
}

func HandleTextProperties(textfile string) ([]string,error) {	//返回切片和error信息
	 fmt.Println("properties : "+textfile )
    file, err := os.Open(textfile)
    	var infoSlice []string =make([]string,0)
    if err != nil {
        log.Printf("Cannot open text file: %s, err: [%v]", textfile, err)
        return infoSlice,err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
    
        line := scanner.Text() 
	    
	     //过滤普通用户用户名
	     if !strings.Contains(line,"--"){
	     	//加载欲操作
		     if strings.Contains(line,"Database"){
		     	infoSlice = append(infoSlice,""+line)
		     }
		     if strings.Contains(line,"ClusterDb"){
		     	infoSlice = append(infoSlice,""+line)
		     }
		     if strings.Contains(line,"FullBackup"){
		     	infoSlice = append(infoSlice,""+line)
		     }
		     if strings.Contains(line,"BackupOrRestore"){
		     	infoSlice = append(infoSlice,""+line)
		     }
		     if strings.Contains(line,"DBUrl"){
		     	infoSlice = append(infoSlice,""+line)
		     }
		     if strings.Contains(line,"BkUsr"){
		     	infoSlice = append(infoSlice,""+line)
		     }
		     if strings.Contains(line,"DBPort"){
		     	infoSlice = append(infoSlice,""+line)
		     }
		     if strings.Contains(line,"BkPass"){
		     	infoSlice = append(infoSlice,""+line)
		     }
		     if strings.Contains(line,"BkPath"){
		     	infoSlice = append(infoSlice,""+line)
		     }
		     if strings.Contains(line,"BackupPth"){
		     	infoSlice = append(infoSlice,""+line)
		     }
		     if strings.Contains(line,"RestorePth"){
		     	infoSlice = append(infoSlice,""+line)
		     }
		     if strings.Contains(line,"SshUsr"){
		     	infoSlice = append(infoSlice,""+line)
		     }
		     if strings.Contains(line,"SshPass"){
		     	infoSlice = append(infoSlice,""+line)
		     }
		      if strings.Contains(line,"DBPth"){
		     	infoSlice = append(infoSlice,""+line)
		     }
		     if strings.Contains(line,"DBSock"){
		     	infoSlice = append(infoSlice,""+line)
		     }
			 if strings.Contains(line,"DBConf"){
		     	infoSlice = append(infoSlice,""+line)
		     }
		     if strings.Contains(line,"StartTime"){
		     	infoSlice = append(infoSlice,""+line)
		     }
		     if strings.Contains(line,"CheckTime"){
		     	infoSlice = append(infoSlice,""+line)
		     }
	     
	     //user_name := strings.Split(line,"=")
//	     infoSlice = append(infoSlice,""+line)
		}
	    
		   
	 }

    if err := scanner.Err(); err != nil {
        log.Printf("Cannot scanner text file: %s, err: [%v]", textfile, err)
        return infoSlice,err
    }

	return infoSlice,nil
}