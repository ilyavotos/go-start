package main

import "fmt"

type Vehicle interface {
	move()
}
type Car struct {
	name string
}
type Ship struct {
	name string
}

func (c Car) move() {
	fmt.Println(c.name)
}

func (s Ship) move() {
	fmt.Println(s.name)
}

func drive(v Vehicle) {
	v.move()
}

func main() {
	car := Car{"Volvo"}
	ship := Ship{"Admiral"}

	drive(car)
	drive(ship)
}
