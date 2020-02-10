package main

import "fmt"

func main(){

	add,sub :=additions_and_substraction(5,2)
	fmt.Print("Addtion is : ",add ,"  and substraction is : ",sub)
	fmt.Println()

	// only get the addition of these two numbers..

	add,_ = additions_and_substraction(7,3)
	fmt.Print("Additions is : ",add)
}

func additions_and_substraction(a,b int)(int,int){
	
	return (a+b),(a-b)
}