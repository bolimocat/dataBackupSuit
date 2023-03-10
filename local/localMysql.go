package local

import (
	"fmt"
	"io"
	"log"
	"os/exec"
	"strings"
)

func asyncLog2(reader io.ReadCloser) error {
	cache := "" //缓存不足一行的日志信息
	buf := make([]byte, 1024)
	for {
		num, err := reader.Read(buf)
		if err != nil && err!=io.EOF{
			return err
		}
		if num > 0 {
			b := buf[:num]
			s := strings.Split(string(b), "\n")
			line := strings.Join(s[:len(s)-1], "\n") //取出整行的日志
			fmt.Printf("%s%s\n", cache, line)
			cache = s[len(s)-1]
		}
	}
	return nil
}

func Localmysql(dbpth string,dbuser string,dbpass string,dbhost string,dbport string,sqlcmd string) error {
	
	 cmd := exec.Command("sh", "-c", dbpth+"/bin/mysql -u'"+dbuser+"' -p'"+dbpass+"' -h'"+dbhost+"' -P'"+dbport+"' -e ''"+sqlcmd+"''")
		stdout, _ := cmd.StdoutPipe()
		stderr, _ := cmd.StderrPipe()

		if err := cmd.Start(); err != nil {
		log.Printf("Error starting command: %s......", err.Error())
		return err
	}

	go asyncLog2(stdout)
	go asyncLog2(stderr)

	if err := cmd.Wait(); err != nil {
		log.Printf("Error waiting for command execution: %s......", err.Error())
		return err
	}

	return nil
    
}

