package main

import "fmt"

func nojam11050() {
	var a int
	var b int

	fmt.Scanln(&a)
	fmt.Scanln(&b)
	var FacA = 1
	var FacB = 1
	var FacC = 1
	for i := 1; i <= a; i++ {
		FacA *= i
	}
	for j := 1; j <= b; j++ {
		FacB *= j
	}
	for k := 1; k <= (a - b); k++ {
		FacC *= k
	}
	var resultA int
	fmt.Println(FacA)
	fmt.Println(FacB)
	fmt.Println(FacC)
	resultA = FacA / (FacB * FacC)
	fmt.Println(resultA)
}

func main() {
	nojam11050()
}
