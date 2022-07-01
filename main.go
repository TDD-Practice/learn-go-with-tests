package main

import "fmt"

const engHello = "Hello, "
const espHello = "Hola, "


func helloWorld() string {
	return "Hola Matheus!"
}


func helloYou(name string) string{
	return fmt.Sprintf("%s%s!", engHello, name)
}


func sayHello(name string, lang string) string {
	
	getGreet := func (lang string) string {
		greet := ""

		if lang == "ENG" { greet = engHello}
		if lang == "ESP" { greet = espHello}
		
		return greet
	}

	if name == "" {name = "person"}
	
	if lang  == "" {lang = "ENG"}

	greet := getGreet(lang)
	return fmt.Sprintf("%s%s!", greet, name)
}