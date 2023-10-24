package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

type Action struct {
	Human
}

func main() {
	a := Action{Human{
		Name: "John",
		Age:  40,
	}}
	a.Eat()
	a.Say()
	a.Fight()
	a.Poo()
}

func (h *Human) Say() {
	fmt.Println("Human speaking!")
}

func (h *Human) Eat() {
	fmt.Println("Human eating!")
}

func (h *Human) Poo() {
	fmt.Println("Human thinking about eternity!")
}

func (h *Action) Fight() {
	fmt.Println("Somebody fighting!")
}
