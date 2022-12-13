package main

import (
	"fmt"
	"net"
	"os"
	"qdb/src/commons"
	"qdb/src/kv"
	"strings"
)

func main() {
	args := os.Args
	kvStore, err := kv.OpenKvStore(commons.DATA_BASE_PATH)
	server, err := net.Listen("tcp", args[1])
	if err != nil {
		fmt.Println(err)
	}
	defer server.Close()
	for {
		conn, err := server.Accept()
		if err != nil {
			fmt.Println(err)
		}
		//defer conn.Close()
		go process(conn, kvStore)
	}
}
func process(conn net.Conn, kvStore *kv.KvStore) {
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		n, _ := conn.Read(buf)
		cmd := string(buf[:n])
		cmds := strings.Split(cmd, " ")
		switch cmds[0] {
		case "get":
			val, err := kvStore.Get(cmds[1])
			if err != nil {
				conn.Write([]byte(err.Error()))
			} else {
				conn.Write([]byte(val))
			}
		case "set":
			err := kvStore.Set(cmds[1], cmds[2])
			if err != nil {
				conn.Write([]byte(err.Error()))
			} else {
				conn.Write([]byte("set success."))
			}
		case "remove":
			err := kvStore.Remove(cmds[1])
			if err != nil {
				conn.Write([]byte(err.Error()))
			} else {
				conn.Write([]byte("remove success."))
			}
		}
	}
}
