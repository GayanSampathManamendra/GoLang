package main

import "fmt"

func add1(a int , b int)int{
	return (a+b)
}

func main(){
	
	ans_add1 := add1(2,3)
	fmt.Println(ans_add1)

	ans_add2 := add2(5,6,6)
	fmt.Println(ans_add2)
	
}

func add2(a,b,c int)int{
	return (a+b+c)
}
