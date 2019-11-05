package main

import "fmt"

var container = []string{"zero", "one", "two"}

func main() {
	container := map[int]string{0: "zero1", 1: "one1", 2: "two1"}
	fmt.Printf("The element is %q.\n", container[1])
}
