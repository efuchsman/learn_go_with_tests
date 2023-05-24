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
