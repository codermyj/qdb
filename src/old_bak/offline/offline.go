package main

import (
	"bufio"
	"fmt"
	"os"
	api2 "qfs/src/old_bak/api"
	"strings"
)

func main() {
	keyDir := api2.LoadData()
	var pos int64 = 0
	for {
		var data api2.TheData
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
			api2.SetData(data, keyDir, &pos)
			fmt.Println("修改成功")
		//case "rm":
		//	api.RmData(data, allData)
		//	fmt.Println("删除成功")
		case "get":
			find, ok := api2.GetData(data.Key, keyDir)
			if ok == false {
				fmt.Println("无记录")
			} else {
				fmt.Println("查询结果", find)
			}
		}
	}
}
