package main

import (
	"log"
	"os"
	"fmt"
	"path/filepath"
	"strings"
	"cLinks"
)

//获取可执行文件路径
func GetFileDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
	//return dir
}

func main() {
	//创建命令集合
	linuxCmdMap := make(map[string]string)
	linuxCmdMap["查看cpu数量"] = "cat /proc/cpuinfo|grep processor|wc -l"
	linuxCmdMap["查看内存大小"] = "cat /proc/meminfo|grep MemTotal|awk '{print $2}'"
	linuxCmdMap["查看操作系统"] = "cat /etc/redhat-release"
	//balance := [...]string{"cat /proc/cpuinfo|grep processor|wc -l", "cat /proc/meminfo|grep MemTotal|awk '{print $2}'", "cat /etc/redhat-release"}
	//for i := 0; i<len(balance); i++{
	//	//fmt.Println(i)
	//	shellc := balance[i]
	//	fmt.Println(shellc)
	//	cLinks.ShellCommand(shellc)
	//}
	for linuxcmdinfo,linuxcmd := range linuxCmdMap  {
		fmt.Println(linuxcmdinfo)
		cLinks.ShellCommand(linuxcmd)
	}
	fmt.Println("可执文件行路径所在位置" + GetFileDirectory())
	//cLinks.ShellCommand("cat /etc/hosts")
}