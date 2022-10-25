package main

import "fmt"

func main() {
	x := make(map[string]string)
	x["xx"] = "dd"
	fmt.Println(x)
	modify(x)
	fmt.Println(x)
}

func modify(m map[string]string) {
	m["aa"] = "bb"
}
