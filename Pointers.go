package main

import "fmt"

func main(){
	
	value:=20
	
	// assign the memory location to the poiner ptr1..
	ptr1:=&value
	
	fmt.Println(value)
	
	// print the momory location of the poiner..
	fmt.Println(ptr1)
	
	// print the vale of the poiner
	fmt.Println(*ptr1)
	
	// change the poiner value
	*ptr1 = 50
	fmt.Println(*ptr1)
	
	// then change the value	
	fmt.Println("new value is : ",value)
	
	// change the value
	value = 100
	fmt.Println("now value is : ", value)
	
	// then change the poiner value..
	fmt.Println("New Poiner value is  : ", *ptr1)
	
	// print the poiner memory location 
	fmt.Println("Poiner memory location : ",ptr1)
}