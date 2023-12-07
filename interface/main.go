package main

import (
	"fmt"
	"math"
)

type ShapeDesc interface {
	Area()	float64
	Perimeter() float64
}

type rectangle struct {
	H, W float64
}

type circle struct {
	R float64
}

func (r *rectangle) Area() float64 {
	return r.H * r.W
}

func (r *rectangle) Perimeter() float64 {
	return 2 * (r.H + r.W)
}

func (c *circle) Area() float64 {
	return c.R * c.R * math.Pi
}

func (c *circle) Perimeter() float64 {
	return 2 * c.R * math.Pi
}

func Desc(s ShapeDesc)  {
	_, ok := s.(*circle)
	if ok {
		fmt.Println("This is circle.")
	}

	_, ok = s.(*rectangle)
	if ok {
		fmt.Println("This is retangle")
	}
	fmt.Println("Area: ", s.Area())
	fmt.Println("Perimeter: ", s.Perimeter())
}

func main()  {
	var s1, s2 ShapeDesc
	s1 = &rectangle{H:2, W:3}
	s2 = &circle{R:3}

	Desc(s1)
	Desc(s2)
}