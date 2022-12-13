package main

import (
	"bufio"
	"fmt"
	"os"
	"qdb/src/commons"
	"qdb/src/kv"
	"strings"
)

func main() {
	kvStore, err := kv.OpenKvStore(commons.DATA_BASE_PATH)
	if err != nil {
		return
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		commons.PrintPrompt()
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
