package goHello

import "fmt"

// SayHello Just say hello
func SayHello(name string) {
	fmt.Println("Hello", name)
}

func SayBye(name string) {
	fmt.Println("bye", name)
}

func Sum(a, b int) int {
	return a + b
}
