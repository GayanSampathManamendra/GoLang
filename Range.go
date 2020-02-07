package main

import "fmt"

func main(){
	
	a := []int{1,2,3,4,5}
	
	for i,c:=range a{
		fmt.Print(a[i], " position value is : ",c)
		fmt.Println()
	}

	// iterate the map ..(key and value pair)

	b := map[int]int{1:2,2:4,3:6}
	
	for k,v := range b {
		fmt.Print("key is :",k , "  and value is : ",v)
		fmt.Println()	
	}

	// range over the key ..
	for k := range b{
		fmt.Println("key ",k)
	}
}