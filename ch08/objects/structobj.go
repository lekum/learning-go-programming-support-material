package main

import "fmt"

type fuel int

const (
	GASOLINE fuel = iota
	BIO
	ELECTRIC
	JET
)

type vehicle struct {
	make  string
	model string
}

type engine struct {
	fuel   fuel
	thrust int
}

func (e *engine) start() {
	fmt.Println("Engine started.")
}

type truck struct {
	vehicle
	engine
	axels  int
	wheels int
	class  int
}

func (t *truck) drive() {
	fmt.Printf("Truck %s %s, on the go!\n", t.make, t.model)
}

type plane struct {
	vehicle
	engine
	engineCount int
	fixedWings  bool
	maxAltitude int
}

func (p *plane) fly() {
	fmt.Printf(
		"Aircraft %s %s clear for takeoff!\n",
		p.make,
		p.model,
	)
}

func main() {
	p := plane{
		vehicle{make: "Airbus", model: "A380"},
		engine{fuel: fuel(1), thrust: 3},
		2,
		true,
		300,
	}
	p.fly()
	p.start()
}
