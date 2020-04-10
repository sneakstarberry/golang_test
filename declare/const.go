package main

import "fmt"

const (
	sunday   = iota // 0
	monday          // 1
	tuesday         // 2
	thursday        // 3
	friday          // 4
	saturday        // 5
)

func main() {
	fmt.Println(sunday)
	fmt.Println(monday)
	fmt.Println(tuesday)
	fmt.Println(thursday)
	fmt.Println(friday)
	fmt.Println(saturday)
}
