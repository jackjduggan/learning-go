// firstly, define a package. Must be done in all Go files.
package main

import "fmt" // fmt is a standard module, almost always imported. Allows for printing things out to the console.

func main() {

	// variables - var {name} {type} = {value}
	var name string = "Hello Jack"
	//var number uint = 1000 // declaring a value and not using it produces an error in Go.
	number := 6          // walrus operator, fasted way to assign a variable
	boolean_var := false // booleans must be lowercase
	// all types have a default value if you don't give them a value.

	fmt.Println("Hello World") // single quotes doesn't work for strings in Go. You use '' for char instead.
	fmt.Println(name, boolean_var)
	fmt.Printf("%T", number) // %T displays the type
}
