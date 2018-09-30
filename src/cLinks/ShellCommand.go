package cLinks
import (
	"fmt"
	"go-gypsy-master/yaml"
	"os"
	"log"
	"strings"
)
//执行命令方法
func ShellCommand(smd string){
	fmt.Println("你输入的命令是：" + smd)
	config, err := yaml.ReadFile( "src/conf/hostInfo.yaml")
	if err != nil {
		fmt.Println(err)
	}
	user, _ := config.Get("user")
	password, _ := config.Get("password")
	hosts, _ := config.Get("host")
	hostsz := strings.Split(hosts,",")
	//fmt.Println(hostsz)
	//port, _ := config.GetInt("port")
	for i:=0;i<len(hostsz);i++{
		host := hostsz[i]
		//fmt.Printf("%s\n",host)
		//host := strconv.Itoa(hosti)
		fmt.Printf("查询主机%s\n",host)
		session, err := Connect(user, password, host, 22)
		if err != nil {
			log.Fatal(err)
		}
		defer session.Close()
		//修改一下输出
		session.Stdout = os.Stdout
		session.Stderr = os.Stderr
		session.Run(smd)
		//return kaka
	}
}
