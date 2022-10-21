package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	m := make(map[string]string)
	m["aa"] = "bb"
	m["cc"] = "dd"
	//fmt.Printf("%v", m)
	msl, _ := json.Marshal(m)
	s := string(msl)
	fmt.Printf("%v, %v", msl, s)
}
