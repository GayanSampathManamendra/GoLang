package main

import "fmt"

type student struct{
	first_name string
	last_name string 
	age int

}

func main(){

	// struct object can be assign with property and value..
	fmt.Println(student{first_name:"gayan",last_name:"sampath",age:25})

	// and also can be assign without property (only value)
	fmt.Println("Gayan","Sampath",23)
	
	// when values are changed according to the property order , values will assign to wrong property..
	fmt.Println("Sampath",25,"Gayan")
	
	// print the property of the particular struct object
	std1 := student{"johon", "smith",41}
	fmt.Println("student1 first name is : ",std1.first_name)
	
	// assign the std1 object to a pointer which is called std1_ptr,
	std1_ptr := &std1
	fmt.Println(std1_ptr.first_name)
	
	// change first name value using pointer..
	std1_ptr.first_name= "Nick"
	fmt.Println(std1_ptr.first_name)
	
	// std1.first_name value also will changed
	fmt.Println(std1.first_name)
}