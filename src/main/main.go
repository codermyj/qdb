package main

import (
	"bufio"
	"fmt"
	"os"
	"qfs/src/ops"
	"strings"
)

func main() {
	for {

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		cmd, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("发生错误, %v", err)
		}
		// fmt.Println("test-----", cmd, len(cmd))
		cmd = cmd[0 : len(cmd)-2]
		cmds := strings.Split(cmd, " ")
		op := cmds[0]
		data := ""
		if len(cmds) > 1 {
			data = cmd[len(op)+1:]
		}
		//fmt.Println(cmds)
		switch op {
		case "add":
			ops.SetData(data)
			fmt.Println("添加一行数据成功")
		case "rm":
			fmt.Println("待实现...")
		case "find":
			find := ops.GetData(data)
			fmt.Println("检索结果", find)

		}

	}
}
