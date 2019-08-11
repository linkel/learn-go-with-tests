package main

import "fmt"

// Hello is a function that takes in a string and concatenates it with hello
func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Println(Hello("world"))
}
