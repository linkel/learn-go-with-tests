package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const cantonese = "Cantonese"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const cantoneseHelloPrefix = "Nei hou, "

// Hello is a function that takes in a string and concatenates it with hello
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case cantonese:
		prefix = cantoneseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("", ""))
	fmt.Println(Hello("Dude", "Cantonese"))
}
