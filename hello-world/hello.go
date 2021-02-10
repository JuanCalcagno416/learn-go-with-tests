package main

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"
)

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchhHelloPrefix = "Bonjour, "
)

func main() {
	fmt.Println(Hello("Nacho", "Spanish"))
}

// Hello will hello you and the world :troll:
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return GreetingPrefix(language) + name
}

// GreetingPrefix will return a prefix depending on the provided language
func GreetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchhHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
