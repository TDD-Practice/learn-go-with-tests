package main

import (
	"fmt"
	"io"
	"os"
)

func InjectableGreeter(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {
	InjectableGreeter(os.Stdout, "Matheus")
}
