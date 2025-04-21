package main

import "fmt"

const (
	englishHelloPrefix = "Hello,"
	SpanishHelloPrefix = "Hola,"
	FrenchHelloPrefix  = "Bonjour,"
	spanish            = "spanish"
	french             = "French"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	return RetrunGreetingPrefix(language) + name

}

func RetrunGreetingPrefix(language string) (prefix string) {

	switch language {
	case spanish:
		prefix = SpanishHelloPrefix
	case french:
		prefix = FrenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return prefix
}

func main() {
	fmt.Println(Hello("chris1", "english"))
}
