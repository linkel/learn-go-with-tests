package main

import "fmt"

const englishHelloPrefix = "Hello, "

// Hello is a function that takes in a string and concatenates it with hello
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
