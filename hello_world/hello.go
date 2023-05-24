package hello_world

func Hello() string {
	return "Hello, World"
}

const englishHelloPrefix = "Hello, "

func HelloWithName(name string) string {

	return englishHelloPrefix + name
}
