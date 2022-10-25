package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	listen, err := net.Listen("tcp", "172.16.39.228:20000")
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
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
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
		inputReader := bufio.NewReader(os.Stdin)
		s, _ := inputReader.ReadString('\n')
		t := strings.Trim(s, "\r\n")
		conn.Write([]byte(t))
	}
}
