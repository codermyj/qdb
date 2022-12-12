package main

import (
	"fmt"
)

func main() {
	b := make([]byte, 0, 4)

	fmt.Println(b)

	bytes := append(b, []byte{1, 2}...)
	fmt.Println(bytes)
}
