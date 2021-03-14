package main

import (
	"fmt"
)

var x = string("fart")

func main() {
	switch x {
	case "foo":
		fmt.Println("It's Foo")
	case "bar":
		fmt.Println("It's bar")
	default:
		fmt.Println("It's not foobar")
	}

}
