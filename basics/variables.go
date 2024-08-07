package main

import "fmt"

func main() {

	var a = "initial" // string
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int // zero valued, defaults to 0
	fmt.Println(e)

	f := "apple" // shorthand for declaring and initialising a variable (var f string = "apple")
	fmt.Println(f)

	a1 := "initial"
	b1, c1 := 1, 2
	d1 := true

	fmt.Println(a1, b1, c1, d1)

}
