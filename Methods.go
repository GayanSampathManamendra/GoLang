package main

import "fmt"

func main(){

	r := rect{width:5,heigth:10}
	
	fmt.Println("Area is : ",r.area())
	
	// get the perim using pointers..
	
	rec_ptr:=&r
	
	fmt.Println("Perim of the rectrangle using pointers : ",rec_ptr.perim())
}

type rect struct{

	width, heigth int
}

func (r, rect) area()int{
	return r.width*r.heigth
}

func (r, *rect) perim()int{
	return 2*r.width*2*r.heigth
}
