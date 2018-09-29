package cLinks
import (
	"fmt"
	"go-gypsy-master/yaml"
	"os"
	"log"
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
	host, _ := config.Get("host")
	//port, _ := config.GetInt("port")
	//fmt.Printf("%s\n",user)
	session, err := Connect(user, password, host, 22)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Run(smd)
	//return kaka
}
