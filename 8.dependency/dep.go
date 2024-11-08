package dependency

import (
	"fmt"
	"io"
)

// func Greet(w *bytes.Buffer, name string) {
// 	fmt.Fprintf(w, "Hello, %s", name)
// }

// io.Writer is an interface
// type Writer interface {
// 	Write(p []byte) (n int, err error)
// }

// fmt.Fprintf internally calls a Write() func which statisfies the io.Writer interface
// Example of a custom type which "implements" the io.Writer interface

// type MyWriter struct{}

// func (mw MyWriter) Write(p []byte) (n int, err error) {
//     fmt.Printf("Writing %d bytes: %s\n", len(p), p)
//     return len(p), nil
// }

func Greet(w io.Writer, name string) {
	fmt.Fprintf(w, "Hello, %s", name)
}
