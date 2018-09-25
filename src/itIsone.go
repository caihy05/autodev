package main

import (
	"cLinks"
	"log"
	"os"
	"go-gypsy-master/yaml"
	"fmt"
	"path/filepath"
	"strings"
)

func GetFileDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
	//return dir
}

func main() {
	fmt.Println(GetFileDirectory())
	config, err := yaml.ReadFile("E:/learn/go/autodev/src/conf/hostInfo.yaml")
	if err != nil {
		fmt.Println(err)
	}
	user, _ := config.Get("user")
	password, _ := config.Get("password")
	host, _ := config.Get("host")
	//port, _ := config.GetInt("port")
	fmt.Printf("%s\n",user)
	session, err := cLinks.Connect(user, password, host, 22)
	if err != nil {
	log.Fatal(err)
	}
	defer session.Close()
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Run("lsblk")
}