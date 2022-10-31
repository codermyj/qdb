package main

import (
	"fmt"
	"net"
	"qfs/src/ops"
	"strings"
)

func main() {
	config := getConfig()
	allData := ops.LoadData()
	addr := config["addr"] + ":" + config["port"]
	listen, err := net.Listen("tcp", addr)
	fmt.Printf("服务端: %T=======\n", listen)
	if err != nil {
		fmt.Println("监听失败，err: ", err)
		return
	}
	for {
		conn, err := listen.Accept()
		fmt.Println("当前建立了tcp连接")
		if err != nil {
			fmt.Println("建立连接失败, err: ", err)
			continue
		}
		go process(conn, allData)
	}
}

func process(conn net.Conn, allData map[string]string) {
	defer conn.Close()
	var data ops.Data
	fmt.Printf("服务端：%T\n", conn)
	for {
		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("从客户端读取数据发生错误", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("服务端收到客户端发来的数据", recvStr)

		cmds := strings.Split(recvStr, " ")
		op := cmds[0]
		data.Key = cmds[1]
		data.KeySize = len(data.Key)
		if len(cmds) > 2 {
			data.Value = cmds[2]
			data.ValueSize = len(data.Value)
		}

		var res string

		switch op {
		case "set":
			ops.SetData(data, allData)
			//fmt.Println("修改成功")
			res = "修改成功"
		case "rm":
			ops.RmData(data, allData)
			//fmt.Println("删除成功")
			res = "删除成功"
		case "get":
			find, ok := ops.GetData(data.Key, allData)
			if ok == false {
				//fmt.Println("无记录")
				res = "无记录"
			} else {
				//fmt.Println("查询结果", find)
				res = "查询结果: " + find
			}
		}
		conn.Write([]byte(res))
	}
}
