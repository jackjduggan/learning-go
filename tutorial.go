// firstly, define a package. Must be done in all Go files.
package main

import "fmt" // fmt is a standard module, almost always imported. Allows for printing things out to the console.

func main() {
	fmt.Println("hello world") // single quotes doesn't work for strings in Go. You use '' for char instead.
}
