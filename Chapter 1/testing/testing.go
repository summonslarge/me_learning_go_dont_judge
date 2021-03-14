package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

type Dog struct{}

type Fail struct {
	Number int
}

func (d *Dog) SayHello() {
	fmt.Println("Woof!")
}

func (p *Person) SayHello() {
	fmt.Println("Hello, ", p.Name)
}

type Friend interface {
	SayHello()
}

func Greet(f Friend) {
	f.SayHello()
}

func main() {
	var guy = new(Person)
	guy.Name = "Dave"
	Greet(guy)

	var dog = new(Dog)
	Greet(dog)

}
