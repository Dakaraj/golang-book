package main

import "fmt"

func main() {
	x := "Hello, World!"
	fmt.Println(x)
	f(x)
}

func f(str string) {
	fmt.Println(str)
}
