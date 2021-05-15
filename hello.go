package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjur, "
const polishHelloPrefix = "Witam, "
const polish = "Polish"
const french = "French"
const spanish = "Spanish"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case polish:
		prefix = polishHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("", "world"))
}
