package main

import "fmt"

// Polymorphism

type Greeter interface {
	Greet() string
}

type Person struct {
	Name string
}

func GreetSomeone(g Greeter) {
	fmt.Println(g.Greet())
}

func (p *Person) Greet() string {
	return fmt.Sprintf("Hello! %s.", p.Name)
}

func main() {
	person := &Person{Name: "Alice"}
	GreetSomeone(person)
	person2 := &Person{Name: "Alexandr"}
	GreetSomeone(person2)
}
