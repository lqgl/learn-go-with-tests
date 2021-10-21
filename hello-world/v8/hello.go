package main

import (
	"fmt"

	"golang.org/x/text/language"
)

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

	// if lanuage == spanish {
	// 	return spanishHelloPrefix + name
	// }

	// if lanuage == french {
	// 	return frenchHelloPrefix + name
	// }

	// return englishHelloPrefix + name

	return greetingPrefix(lanuage) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language{
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
