package main

import "fmt"

var block = "package"
var i = 1

func main() {
	i := "sss"
	fmt.Println(i)
	block := "function"
	{
		block := "inner"
		fmt.Printf("The block is %s.\n", block)
	}
	fmt.Printf("The block is %s.\n", block)
	print()
}
func print() {
	fmt.Println(block)
}
