package main

import (
	"bufio"
	"fmt"
	"os"
	"qfs/src/ops"
	"strings"
)

func main() {
	allData := ops.LoadData()
	for {
		var data ops.Data
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("qfs> ")
		cmd, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("发生错误, %v", err)
		}
		// fmt.Println("test-----", cmd, len(cmd))
		cmd = cmd[0 : len(cmd)-2]
		cmds := strings.Split(cmd, " ")
		op := cmds[0]
		data.Key = cmds[1]
		data.KeySize = len(data.Key)
		if len(cmds) > 2 {
			data.Value = cmds[2]
			data.ValueSize = len(data.Value)
		}

		switch op {
		case "set":
			ops.SetData(data, allData)
			fmt.Println("添加一行数据成功")
		case "rm":
			ops.RmData(data, allData)
			fmt.Println("删除记录...")
		case "get":
			find := ops.GetData(data.Key, allData)
			fmt.Println("检索结果", find)
		}
	}
}
