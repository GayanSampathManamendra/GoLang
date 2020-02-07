package main

import "fmt"

func main(){

	// create a slice..
	s := make([]int,3)
	fmt.Println(s)

	s[0]=5
	s[1]=3
	s[2]=4

	fmt.Println(s)
	fmt.Println("Slice Length is :",len(s))

	// append function is unique to slice..
	s=append(s,8,9,3,4)
	fmt.Println("New slice is :",s)

	// get the sub array using the slice..
	fmt.Println(s[1:2])

	fmt.Println(s[:5])

	fmt.Println(s[2:])


	// define a new slice and asssign it to previous one
	x:=s
	fmt.Println(x)

	// change the first value of the newly initilized slice x
	x[0]=100
	fmt.Println(x)

	
	// then change the fist value of previous defined slice
	fmt.Println(x)

	// to prevent that issue..copy the previous slice as a new slice.
	// create the new slice and assigned to x 
	x = make([]int , len(s))
	copy(x,s)
	
	// print the x slice
	fmt.Println(x)

	// print the s slice
	s[0]=5
	fmt.Println(s)


	// 2-D slice.
	d := make([][]int , 3)
	fmt.Println(d)	

	for i:=0;i<3;i++{
	    inner_length := i +1
	    d[i] = make([]int,inner_length)
	}
	fmt.Println(d)
}