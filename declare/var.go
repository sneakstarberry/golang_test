package main

import "fmt"

var (
	num   int
	float float32
	exam  string
)

const (
	sunday   = iota // 0
	monday          // 1
	tuesday         // 2
	thursday        // 3
	friday          // 4
	saturday        // 5
)

func main() {
	num = 10
	float = 5.23
	exam = "Hello World"

	var num2 = 20

	fmt.Println("int", num)
	fmt.Println("float", float)
	fmt.Println("string", exam)

	fmt.Println("int", num2)
}
