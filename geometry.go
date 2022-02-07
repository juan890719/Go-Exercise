package main

import "fmt"

type geometry interface {
	area() float64
}

type retangle struct {
	width  float64
	height float64
}

func (r *retangle) area() float64 {
	return r.width * r.height
}

type circle struct {
	radius float64
}

func (c *circle) area() float64 {
	return c.radius * c.radius * 3.1416
}

func printRetangleArea() {
	r1 := &retangle{
		width:  23.4,
		height: 42.8,
	}

	c1 := &circle{
		radius: 2,
	}

	fmt.Println(r1.area())
	fmt.Println(c1.area())
}
