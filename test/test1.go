package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.OpenFile("./test.txt", os.O_APPEND, 0666)
	defer file.Close()
	var s string
	//defer
	for {
		fmt.Printf("> ")
		reader := bufio.NewReader(os.Stdin)
		s1, _ := reader.ReadString('\n')
		s += s1
	}
	file.WriteString(s)
}
