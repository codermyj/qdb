package main

import "fmt"

func main() {
	slice1 := []int{999: 999}
	fmt.Printf("slice1: len=%v, cap=%v\n", len(slice1), cap(slice1))
	slice2 := slice1[1:3]
	fmt.Printf("slice2: len=%v, cap=%v\n", len(slice2), cap(slice2))

	slice1 = append(slice1, 5)
	fmt.Printf("slice1: len=%v, cap=%v\n", len(slice1), cap(slice1))
	fmt.Printf("slice2: len=%v, cap=%v\n", len(slice2), cap(slice2))
}
