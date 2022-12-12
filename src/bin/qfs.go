package main

import (
	"bufio"
	"fmt"
	"os"
	"qfs/src/kv"
	"strings"
)

const DATA_BASE_PATH = "./data/"

func printPrompt() {
	fmt.Printf("%v> ", kv.STORAGE_FILE_PREFIX)
}

func main() {
	kvStore, err := kv.OpenKvStore(DATA_BASE_PATH)
	if err != nil {
		return
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		printPrompt()
		line, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println(err)
			continue
		}
		cmd := string(line)
		cmds := strings.Split(cmd, " ")
		switch cmds[0] {
		case "get":
			val, err := kvStore.Get(cmds[1])
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(val)
		case "set":
			err := kvStore.Set(cmds[1], cmds[2])
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("set操作成功")
		case "remove":
			err := kvStore.Remove(cmds[1])
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("删除成功")
		}
	}
}
