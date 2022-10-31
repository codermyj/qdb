package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getConfig() map[string]string {
	file, err := os.Open("./config")
	if err != nil {
		fmt.Println("读取配置文件失败, err:", err)
	}
	defer file.Close()
	config := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		split := strings.Split(str, "=")
		config[split[0]] = split[1]
	}
	fmt.Println(config)
	return config
}
