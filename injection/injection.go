package injection

import (
	"bytes"
	"fmt"
)

func InjectableGreeter(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)

}
