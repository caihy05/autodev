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
	balance := [...]string{"lsblk", "df -h", "free"}
	for i := 0; i<len(balance); i++{
		//fmt.Println(i)
		shellc := balance[i]
		fmt.Println(shellc)
		cLinks.ShellCommand(shellc)
	}
	fmt.Println("路径" + GetFileDirectory())
	//cLinks.ShellCommand("cat /etc/hosts")
}