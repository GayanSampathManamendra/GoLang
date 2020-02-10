package main

import (
		"fmt"
		"math"
)

type geometry interface{
	area() float64
}

type rect stuct{
	width, heigth int
}

type circle stuct{
	radius int
}

func (c circle) area()float64{
	return 2*r*math.Pi*c.radius
}

func (r rect) area()float64{
	return r.width*r.heigth
}


func getArea(g geometry)float64{
	
	fmt.Println("Geometry is : ",g)
	fmt.Println(g.area())
}

func main(){

	r := rect{width:5,heigth:10}
	c := circle{radius:5}
	
	getArea(r)
	getArea(c)
	
}