package main

import (
	"fmt"
)

const english = "English"
const spanish = "Spanish"
const french = "French"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish :
		prefix = spanishHelloPrefix
	case french :
		prefix = frenchHelloPrefix
	default :
		prefix = englishHelloPrefix
	}
	return
}


func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	if language == "" {
		language = english
	}

	return greetingPrefix(language) + name
}


func main(){
	fmt.Println(Hello("Darth Vader", "Spanish"))
}