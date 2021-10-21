package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const spanish = "Spanish"
const french = "French"

// Hello returns a personalised greeting.
func Hello(name string, lanuage string) string {
	if name == "" {
		name = "World"
	}

	if lanuage == spanish {
		return spanishHelloPrefix + name
	}

	if lanuage == french {
		return frenchHelloPrefix + name
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
