package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type poligon interface {
	shape
	perim()
}

type curved interface {
	shape
	circonf()
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

type triangle struct {
	name    string
	a, b, c float64
}

func (t *triangle) area() float64 {
	return 0.5 * (t.a * t.b)
}

func (t *triangle) perim() float64 {
	return t.a + t.b + math.Sqrt((t.a*t.a)+(t.b*t.b))
}

func (t *triangle) String() string {
	return fmt.Sprintf(
		"%s[sides: a=%.2f b=%.2f c=%.2f]",
		t.name, t.a, t.b, t.c,
	)
}

type circle struct {
	name string
	rad  float64
}

func (c *circle) area() float64 {
	return math.Pi * (c.rad * c.rad)
}

func (c *circle) circonf() float64 {
	return 2 * math.Pi * c.rad
}

func shapeInfo(s shape) string {
	return fmt.Sprintf(
		"Area = %.2f",
		s.area(),
	)
}

func main() {
	r := &rect{"Square", 4.0, 4.0}
	fmt.Println(r, "=>", shapeInfo(r))
	t := &triangle{"Right Triangle", 1, 2, 3}
	fmt.Println(t, "=>", shapeInfo(t))
	c := &circle{"Circle", 3}
	fmt.Println(c, "=>", shapeInfo(c))
}
