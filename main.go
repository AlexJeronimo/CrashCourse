package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
	var a int
	a = 14
	fmt.Println(a)
	if a < 0 {
		fmt.Println("a < 0")
	}
	else {fmt.Println("a > 0")}
}
