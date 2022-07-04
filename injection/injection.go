package injection

import "bytes"

func InjectableGreeter(buffer *bytes.Buffer, name string) {
	buffer.WriteString("Hello, ")
	buffer.WriteString(name)
}
