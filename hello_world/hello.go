package hello_world

func Hello() string {
	return "Hello, World"
}

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func HelloWithName(name string) string {

	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func HelloWithLanguage(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "french", "French":
		prefix = frenchHelloPrefix
	case "spanish", "Spanish":
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
