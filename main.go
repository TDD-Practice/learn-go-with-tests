package main

import "fmt"

const engHello = "Hello, "


func helloWorld() string {
	return "Hola Matheus!"
}


func helloYou(name string) string{
	return fmt.Sprintf("%s%s!", engHello, name)
}


func sayHello(name string) string {
	if name == "" {
		name = "person"
	}
	return fmt.Sprintf("%s%s!", engHello, name)
}