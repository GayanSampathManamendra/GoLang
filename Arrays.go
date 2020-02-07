package main

import "fmt"

func main(){
	
	// There are defferent type of array initialization..
		
	// initialize an integer array..
	
	var integer_array[5]int 

	fmt.Println(integer_array)

	// initialize a string array ..

	var str_arr[5]string 
	fmt.Println(str_arr)

	// assign a value the array ..
	
	str_arr[0]="hello"
	str_arr[1]="world"
	fmt.Println("This is new string array :",str_arr)
	fmt.Println("Length of the string array :",len(str_arr))

	// define two dimession array..
	var tda[5][5]int 
	fmt.Println(tda)

	
	for i:=0;i<5;i++{
		for j:=0;j<5;j++{
		    fmt.Print(tda[i][j])
		}
		fmt.Println()
	}

}