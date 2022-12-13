package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"qdb/src/commons"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	in := bufio.NewReader(os.Stdin)

	for {
		commons.PrintPrompt()
		bt, _, err := in.ReadLine()
		if err != nil {
			fmt.Println("Error!!!")
		} else {
			conn.Write(bt)
			bufR := make([]byte, 1024)
			n, _ := conn.Read(bufR)
			fmt.Println(string(bufR[:n]))
		}
	}
}
