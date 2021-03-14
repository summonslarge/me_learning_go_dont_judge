package main

import "errors"

func foo() error {
	return errors.New("Bad thing happened.")
}

func main() {
	if err := foo(); err != nil {
		//Handle Error
	}
}
