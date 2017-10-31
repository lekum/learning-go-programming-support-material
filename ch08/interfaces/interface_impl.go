package main

import "fmt"

type shape interface {
	area() float64
	perim() float64
}

type rect struct {
	name           string
	length, height float64
}

func (r rect) area() float64 {
	return r.length * r.height
}

func (r rect) perim() float64 {
	return 2*r.length + 2*r.height
}

func getPerimeter(s shape) float64 {
	return s.perim()
}

func main() {
	a := rect{"A", 3, 2}
	b := rect{"B", 5, 2.3}
	fmt.Println(getPerimeter(a))
	fmt.Println(getPerimeter(b))
}
