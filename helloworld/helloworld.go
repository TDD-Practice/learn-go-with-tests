package helloworld

import "fmt"

const engHello = "Hello"
const espHello = "Hola"
const fraHello = "Bonjour"

func helloWorld() string {
	return "Hola Matheus!"
}

func helloYou(name string) string {
	return fmt.Sprintf("%s%s!", engHello, name)
}

func sayHello(name string, lang string) string {

	getGreet := func(lang string) (greet string) {
		switch lang {
		case "ENG":
			greet = engHello
		case "ESP":
			greet = espHello
		case "FRA":
			greet = fraHello
		default:
			greet = engHello
		}

		return
	}

	if name == "" {
		name = "person"
	}

	greet := getGreet(lang)
	return fmt.Sprintf("%s, %s!", greet, name)
}
