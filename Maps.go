package main

import "fmt"

func main(){

	// maps in go
	
	a := make(map[string]int)
	
	// string is the key value of the map and int is the value of the map..
	a["first"]=1
	a["second"]=2
	a["third"]=3

	fmt.Println(a)

	// print the specific  value from the map
	
	fmt.Println("Hash map first value is :", a["first"])

	// get the length of the map
	fmt.Println("Hash map length is : ",len(a))

	// delete element form the map
	delete(a,"first")
	fmt.Println(a)

	// another way to define a map..
	
 	b := map[string]int{"forth":4,"fifth":5}
	fmt.Println(b)

	
}