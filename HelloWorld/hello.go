package main

import (
	"fmt"
	"strings"
)

var languagePrefixes = make(map[string]int)

func main() {
	fmt.Println(Hello("Chris", "spanish"))
}

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	if strings.ToLower(language) == "spanish" {
		return "Hola, " + name
	}

	if strings.ToLower(language) == "french" {
		return "Bonjour, " + name
	}

	if strings.ToLower(language) == "german" {
		return "Hallo, " + name
	}

	return "Hello, " + name
}

func getSalutePrefixByLanguage(language string) {

}
