package main

const japanese = "Japanese"
const french = "French"
const englishHelloPrefix = "Hello, "
const japaneseHelloPrefix = "こんにちわ, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case japanese:
		prefix = japaneseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

//func main() {
//	fmt.Println(Hello("world", ""))
//}
