package injection

import (
	"fmt"
	"io"
)

func InjectableGreeter(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
