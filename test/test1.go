package main

import (
	"fmt"
	"github.com/duke-git/lancet/v2/convertor"
)

func main() {
	i1 := "哈哈"

	//b := byte(i1)

	b, _ := convertor.ToBytes(i1)

	fmt.Printf("%v\n", b)

	//fmt.Printf("%v\n", u)
}
