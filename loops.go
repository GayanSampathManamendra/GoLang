package main

import "fmt"

func main(){
	
	// go language only provides "for" loop 		

	i:=0

	fmt.Println("This is first way ...")
	for i<= 10 {
		fmt.Print(i)
		i++
	}
	fmt.Println()

	fmt.Println("This is scecond way ...")	
	for j:=0;j<10;j++{
		fmt.Print(j)
	}
	fmt.Println()

	fmt.Println("This is thired way (Infinite loop )... ")	
	for{
		fmt.Println("Infinite Loop ")
		break
	}
}