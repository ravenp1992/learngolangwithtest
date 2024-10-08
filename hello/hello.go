package hello

import "fmt"

const (
	tagalog = "Tagalog"
	spanish = "Spanish"
	french  = "French"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	tagalogHelloPrefix = "Kumusta, "
)

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
	case tagalog:
		prefix = tagalogHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func Greet(name string) string {
	if name == "" {
		name = "Guest"
	}

	return fmt.Sprintf("Welcome, %s", name)
}

func PrintMyName(name string) string {
	return name
}
